package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
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
var trace bool
var contents string

func main() {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode (default: false).")
	flag.BoolVar(&trace, "trace", false, "Enable trace mode (default: false).")
	flag.StringVar(&contents, "contents", "readme,book-index,author-index,learning-paths", "list of content to generate, accepts comma-separated values")
	flag.Parse()
	contents := strings.Split(contents, ",")

	err := config.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}

	if trace {
		log.Printf("%v\n", config.Cfg)
	}

	// Load raw content from yaml files
	sources.LoadBooks(config.Cfg.Sources.BookData)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Load raw content from yaml files
	sources.LoadLearningPaths(config.Cfg.Sources.LearningPaths)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create content dirs
	if err = os.MkdirAll(config.Cfg.Content.LearningPaths, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	// Create auxiliar structure for easy access to learning paths lpData["apis"].Desc
	lpData := map[string]interface{}{}
	for _, lp := range sources.LearningPaths {
		lpData[string(lp.Ref)] = lp
	}

	// Create auxiliar structure for easy access to books booksData["Building Microservices"].Desc
	booksData := map[string]sources.Book{}
	for _, b := range sources.Books {
		booksData[b.Title] = b
	}

	// Create auxiliar structure to search books by learning path or author seamlessly
	authorsData := map[string][]sources.Book{}
	lpBooksData := map[sources.LearningPathRef][]sources.Book{}
	for _, b := range sources.Books {
		// authorsData["name"] = [{book1}, {book2}, ...]
		for _, a := range b.Authors {
			authorsData[a] = append(authorsData[a], b)
		}

		// lpBooksData["apis"] = [{book1}, {book2}, ...]
		for _, r := range b.LearningPathsRefs {
			lpBooksData[r] = append(lpBooksData[r], b)

			// Sort books by order ascendant and by heavier weight
			sort.SliceStable(lpBooksData[r], func(i, j int) bool {
				if lpBooksData[r][i].Order != lpBooksData[r][j].Order {
					return lpBooksData[r][i].Order < lpBooksData[r][j].Order
				}

				// If order is equal then order by weight
				return lpBooksData[r][i].Weight > lpBooksData[r][j].Weight
			})
		}
	}

	// Create auxiliar structure for easy access to badges badgesData["excellent"] = top
	badgesData := map[string]interface{}{}
	for _, b := range config.Cfg.Badges {
		for _, i := range b.BadgeIcons {
			badgesData[i.Name] = i.Code
		}
	}

	if trace {
		log.Printf("loaded learning paths: %v\n", lpData)
		log.Printf("loaded learning paths books: %v\n", lpBooksData)
		log.Printf("loaded books: %v\n", booksData)
		log.Printf("loaded badges: %v\n", badgesData)
	}

	// Initialize stats
	totalBooks := len(sources.Books)
	totalAuthors := len(authorsData)
	totalLPs := len(sources.LearningPaths)
	booksInLPs := make(map[sources.LearningPathRef]int)
	for lp, books := range lpBooksData {
		booksInLPs[lp] = len(books)
	}
	stats.New(totalBooks, totalAuthors, totalLPs, booksInLPs)

	// Prepare template rendering data
	var data = struct {
		LpData              map[string]interface{}
		LpBooksData         []sources.Book
		BooksData           map[string]sources.Book
		AuthorsData         map[string][]sources.Book
		BadgesData          map[string]interface{}
		BookCovers          string
		LearningPathsFolder string
		BookIndex           string
		AuthorIndex         string
		CurrentLearningPath sources.LearningPath
		Stats               stats.Stats
	}{
		LpData:              lpData,
		BooksData:           booksData,
		AuthorsData:         authorsData,
		BookCovers:          config.Cfg.Sources.BookCovers,
		BadgesData:          badgesData,
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

	if slices.Contains(contents, "learning-paths") {
		for _, lp := range sources.LearningPaths {
			// Render learning paths that are only marked as either stable, new or in-progress
			if lp.Status != "coming-soon" {
				data.CurrentLearningPath = lp
				data.LpBooksData = lpBooksData[lp.Ref]

				if !debug {
					file = filepath.Join(config.Cfg.Content.LearningPaths, fmt.Sprintf("%s.md", lp.Ref))
				}
				log.Printf("rendering learning-path %s in %s", lp.Ref, file)

				if err = content.Render(templates, "learning-path.md.tmpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	log.Println("done.")
}
