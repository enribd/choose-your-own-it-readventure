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
	"github.com/enribd/choose-your-own-it-readventure/internal/sources"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/slices"
)

var debug bool
var contents string
var totalSkippedBooks int

func main() {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode (default: false).")
	flag.StringVar(&contents, "contents", "readme,book-index,author-index,learning-paths", "list of content to generate, accepts comma-separated values")
	flag.Parse()
	contents := strings.Split(contents, ",")

	// Load config
	err := config.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Initialize stats
	stats.New()

	// Load books and learning paths raw content from yaml files
	err = sources.Load(config.Cfg.Sources.BookData, config.Cfg.Sources.LearningPaths)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create content dirs
	if err = os.MkdirAll(config.Cfg.Content.LearningPaths, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	// Create auxiliar structure for easy access to badges badgesData["excellent"] = top
	badgesData := map[string]interface{}{}
	for _, b := range config.Cfg.Badges {
		for _, i := range b.BadgeIcons {
			badgesData[i.Name] = i.Code
		}
	}

	// Prepare template rendering data
	var data = struct {
		LpData              map[string]interface{}
		BooksData           map[string]sources.Book
		AuthorsData         map[string][]sources.Book
		BadgesData          map[string]interface{}
		BookCovers          string
		LearningPathsFolder string
		BookIndex           string
		AuthorIndex         string
		Stats               stats.Stats
		// used only when rendering the learning paths template
		LpBooksData         []sources.Book
		CurrentLearningPath sources.LearningPath
	}{
		LpData:              sources.LearningPathsTmpl,
		BooksData:           sources.Books,
		AuthorsData:         sources.Authors,
		BadgesData:          badgesData,
		BookCovers:          config.Cfg.Sources.BookCovers,
		LearningPathsFolder: config.Cfg.Content.LearningPaths,
		BookIndex:           config.Cfg.Content.BookIndex,
		AuthorIndex:         config.Cfg.Content.AuthorIndex,
		Stats:               stats.Data,
	}

	// Load templates and functions
	templates, err := template.New("base").Funcs(sprig.TxtFuncMap()).ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	// Render content
	file := "stdout" // if in debug mode spit to stdout

	if slices.Contains(contents, "learning-paths") {
		for _, lp := range sources.LearningPaths {
			// Render learning paths that are only marked as either stable, new or in-progress, and have at least 1 book
			if lp.Status != "coming-soon" && len(sources.LearningPathBooks[lp.Ref]) > 0 {
				data.CurrentLearningPath = lp
				data.LpBooksData = sources.LearningPathBooks[lp.Ref]

				if !debug {
					file = filepath.Join(config.Cfg.Content.LearningPaths, fmt.Sprintf("%s.md", lp.Ref))
				}
				log.Printf("rendering '%s' learning path in %s (%d books)", lp.Ref, file, len(sources.LearningPathBooks[lp.Ref]))

				if err = content.Render(templates, "learning-path.md.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Printf("skipping '%s' learning-path status=%s, books=%d", lp.Ref, lp.Status, stats.Data.TotalLearningPathBooks[string(lp.Ref)])
			}
		}
	}

	// Update stats
	if slices.Contains(contents, "book-index") {
		log.Printf("rendering book index in %s", config.Cfg.Content.BookIndex)
		if !debug {
			file = config.Cfg.Content.BookIndex
		}

		if err = content.Render(templates, "book-index.md.tmpl", file, data); err != nil {
			log.Fatalln(err)
		}
	}

	if slices.Contains(contents, "author-index") {
		log.Printf("rendering author index in %s", config.Cfg.Content.AuthorIndex)
		if !debug {
			file = config.Cfg.Content.AuthorIndex
		}

		if err = content.Render(templates, "author-index.md.tmpl", file, data); err != nil {
			log.Fatalln(err)
		}
	}

	if slices.Contains(contents, "readme") {
		log.Printf("rendering readme in %s", config.Cfg.Content.Readme)
		if !debug {
			file = config.Cfg.Content.Readme
		}

		if err = content.Render(templates, "readme.md.tmpl", file, data); err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("learning paths found %d", stats.Data.TotalLearningPaths)
	log.Printf("learning paths skipped %d", stats.Data.TotalSkippedLearningPaths)
	log.Printf("books found %d", stats.Data.TotalBooks)
	log.Printf("books skipped %d", stats.Data.TotalSkippedBooks)
	log.Printf("authors found %d", stats.Data.TotalAuthors)
	log.Println("done.")
}
