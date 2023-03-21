package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/enribd/choose-your-own-it-readventure/config"
	"github.com/enribd/choose-your-own-it-readventure/internal/content"
	"github.com/enribd/choose-your-own-it-readventure/internal/loader"
	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/slices"
)

var debug bool
var formats, contents, mkdocsStripPrefix string

func main() {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode (default: false).")
	flag.StringVar(&formats, "formats", "github,mkdocs", "generate content with format for different hosting providers, accepts comma-separated values")
	flag.StringVar(&contents, "contents", "index,book-index,author-index,learning-paths,badges,about,mentions", "list of content to generate, accepts comma-separated values")
	flag.StringVar(&mkdocsStripPrefix, "mkdocs-strip-path-prefix", "./mkdocs/docs", "remove prefix from path to set browsing routes")
	flag.Parse()
	contents := strings.Split(contents, ",")
	formats := strings.Split(formats, ",")

	// Load config
	err := config.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Initialize stats
	stats.New()

	// Load books and learning paths raw content from yaml files
	err = loader.Load(config.Cfg.Sources.BookData, config.Cfg.Sources.LearningPaths, config.Cfg.Sources.BadgesData)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Loop formats to generate templates
	for _, format := range formats {
		p, err := content.NewProvider(format)
		if err != nil {
			log.Fatalln(err)
			return
		}

		// Arrange content folders
		lpFolder := config.Cfg.Content[p].LearningPaths
		bookIndexFolder := config.Cfg.Content[p].BookIndex
		authorIndexFolder := config.Cfg.Content[p].AuthorIndex
		if p == content.Mkdocs && mkdocsStripPrefix != "" {
			lpFolder = strings.TrimPrefix(lpFolder, mkdocsStripPrefix)
			bookIndexFolder = strings.TrimPrefix(bookIndexFolder, mkdocsStripPrefix)
			authorIndexFolder = strings.TrimPrefix(authorIndexFolder, mkdocsStripPrefix)
		}

		// Create content dirs
		if err = os.MkdirAll(config.Cfg.Content[p].LearningPaths, os.ModePerm); err != nil {
			log.Println(err)
			return
		}

		// Load templates and functions
		providerTmpls := filepath.Join("templates", p.String(), "*")
		funcMap := template.FuncMap{
			"args": content.Args,
		}
		templates, err := template.New("base").Funcs(sprig.TxtFuncMap()).Funcs(funcMap).ParseGlob(providerTmpls)
		if err != nil {
			log.Fatalln(err)
		}

		// Add common templates
		commonTmpls := filepath.Join("templates", "common", "*")
		templates.ParseGlob(commonTmpls)
		if err != nil {
			log.Fatalln(err)
		}

		// Prepare template rendering data
		var data = struct {
			Format              string
			SiteUrl             string
			LpData              map[string]interface{}
			BooksData           map[string]models.Book
			AuthorsData         map[string][]models.Book
			BadgesData          map[string]interface{}
			BookCovers          string
			LearningPathsFolder string
			BookIndex           string
			AuthorIndex         string
			Stats               stats.Stats
			// used only when rendering the learning paths template
			LpBooksData         []models.Book
			CurrentLearningPath models.LearningPath
		}{
			Format:              p.String(),
			SiteUrl:             config.Cfg.SiteUrl,
			LpData:              loader.LearningPathsTmpl,
			BooksData:           loader.Books,
			AuthorsData:         loader.Authors,
			BadgesData:          loader.Badges,
			BookCovers:          config.Cfg.Sources.BookCovers,
			LearningPathsFolder: lpFolder,
			BookIndex:           bookIndexFolder,
			AuthorIndex:         authorIndexFolder,
			Stats:               stats.Data,
		}

		// Render content
		file := "stdout" // if in debug mode spit to stdout

		if slices.Contains(contents, "learning-paths") && config.Cfg.Content[p].LearningPaths != "" {
			for _, lp := range loader.LearningPaths {
				// Render learning paths that are only marked as either stable, new or in-progress, and have at least 1 book
				if lp.Status != "coming-soon" && len(loader.LearningPathBooks[lp.Ref]) > 0 {
					data.CurrentLearningPath = lp
					data.LpBooksData = loader.LearningPathBooks[lp.Ref]

					if !debug {
						file = filepath.Join(config.Cfg.Content[p].LearningPaths, fmt.Sprintf("%s.md", lp.Ref))
					}
					log.Printf("[%s] rendering '%s' learning path in %s (%d books)", p, lp.Ref, file, len(loader.LearningPathBooks[lp.Ref]))

					if err = content.Render(templates, "learning-path.md.tmpl", file, data); err != nil {
						log.Fatalln(err)
					}
				} else {
					log.Printf("[%s] skipping '%s' learning-path status=%s, books=%d", p, lp.Ref, lp.Status, stats.Data.TotalLearningPathBooks[string(lp.Ref)])
				}
			}

			if p == content.Mkdocs {
				file = filepath.Join(config.Cfg.Content[p].LearningPaths, ".pages")
				log.Printf("[%s] rendering .pages in %s", p, file)
				if err = content.Render(templates, "pages_lps.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}

		// Update stats
		if slices.Contains(contents, "book-index") && config.Cfg.Content[p].BookIndex != "" {
			log.Printf("[%s] rendering book index in %s", p, config.Cfg.Content[p].BookIndex)
			if !debug {
				file = config.Cfg.Content[p].BookIndex
			}

			if err = content.Render(templates, "book-index.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}

			if p == content.Mkdocs {
				file = filepath.Join(filepath.Dir(config.Cfg.Content[p].BookIndex), ".pages")
				log.Printf("[%s] rendering .pages in %s", p, file)
				if err = content.Render(templates, "pages_refs.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if slices.Contains(contents, "author-index") && config.Cfg.Content[p].AuthorIndex != "" {
			log.Printf("[%s] rendering author index in %s", p, config.Cfg.Content[p].AuthorIndex)
			if !debug {
				file = config.Cfg.Content[p].AuthorIndex
			}

			if err = content.Render(templates, "author-index.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}

			if p == content.Mkdocs {
				file = filepath.Join(filepath.Dir(config.Cfg.Content[p].AuthorIndex), ".pages")
				log.Printf("[%s] rendering .pages in %s", p, file)
				if err = content.Render(templates, "pages_refs.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if slices.Contains(contents, "index") && config.Cfg.Content[p].Index != "" {
			log.Printf("[%s] rendering index in %s", p, config.Cfg.Content[p].Index)
			if !debug {
				file = config.Cfg.Content[p].Index
			}

			if err = content.Render(templates, "index.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}

			if p == content.Mkdocs {
				// render .pages
				file = filepath.Join(filepath.Dir(config.Cfg.Content[p].Index), ".pages")
				log.Printf("[%s] rendering .pages in %s", p, file)
				if err = content.Render(templates, "pages_index.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}

				// render extra.css
				file = filepath.Join(filepath.Dir(config.Cfg.Content[p].Index), "stylesheets/extra.css")
				log.Printf("[%s] rendering extra.css in %s", p, file)
				if err = content.Render(templates, "extra.css.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if slices.Contains(contents, "badges") && config.Cfg.Content[p].Badges != "" {
			log.Printf("[%s] rendering badges in %s", p, config.Cfg.Content[p].Badges)
			if !debug {
				file = config.Cfg.Content[p].Badges
			}

			if err = content.Render(templates, "badges.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}
		}

		if slices.Contains(contents, "about") && config.Cfg.Content[p].About != "" {
			log.Printf("[%s] rendering about me in %s", p, config.Cfg.Content[p].About)
			if !debug {
				file = config.Cfg.Content[p].About
			}

			if err = content.Render(templates, "about.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}

			if p == content.Mkdocs {
				file = filepath.Join(filepath.Dir(config.Cfg.Content[p].About), ".pages")
				log.Printf("[%s] rendering .pages in %s", p, file)
				if err = content.Render(templates, "pages_more.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}

		if slices.Contains(contents, "mentions") && config.Cfg.Content[p].Mentions != "" {
			log.Printf("[%s] rendering mentions in %s", p, config.Cfg.Content[p].Mentions)
			if !debug {
				file = config.Cfg.Content[p].Mentions
			}

			if err = content.Render(templates, "mentions.md.tmpl", file, data); err != nil {
				log.Fatalln(err)
			}
		}
	}

	log.Printf("learning paths found %d", stats.Data.TotalLearningPaths)
	log.Printf("learning paths skipped %d", stats.Data.TotalSkippedLearningPaths)
	log.Printf("books found %d", stats.Data.TotalBooks)
	log.Printf("books skipped %d", stats.Data.TotalSkippedBooks)
	log.Printf("authors found %d", stats.Data.TotalAuthors)
	log.Printf("done.\n")
}
