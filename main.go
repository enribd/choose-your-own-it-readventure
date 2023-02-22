package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/enribd/choose-your-own-it-readventure/internal/content"
	"github.com/enribd/choose-your-own-it-readventure/internal/sources"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Sources Sources `yaml:"sources"`
	Content Content `yaml:"content"`
	Badges  []Badge `yaml:"badges"`
}

// Location of data used to generate content
type Sources struct {
	BookCovers    string `yaml:"books_covers"`
	BookData      string `yaml:"books_data"`
	LearningPaths string `yaml:"learning_paths_data"`
}

// Destination for generated content
type Content struct {
	Readme        string `yaml:"readme"`
	BookIndex     string `yaml:"book_index"`
	LearningPaths string `yaml:"learning_paths"`
}

type Badge struct {
	Category   BadgeCategory `yaml:"category"`
	BadgeIcons []BadgeIcon   `yaml:"icons"`
}

type BadgeCategory string

type BadgeIcon struct {
	Name string `yaml:"name"`
	Code string `yaml:"code"`
	Desc string `yaml:"desc"`
}

var debug bool
var trace bool
var contents string
var config Config

func main() {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode (default: false).")
	flag.BoolVar(&trace, "trace", false, "Enable trace mode (default: false).")
	flag.StringVar(&contents, "contents", "readme,book-index,learning-paths", "list of content to generate, accepts comma-separated values")
	flag.Parse()
	contents := strings.Split(contents, ",")

	// Read the file
	raw, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Unmarshal the YAML raw content into the struct
	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if trace {
		log.Printf("%v\n", config)
	}

	// Load raw content from yaml files
	sources.LoadBooks(config.Sources.BookData)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Load raw content from yaml files
	sources.LoadLearningPaths(config.Sources.LearningPaths)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create content dirs
	if err = os.MkdirAll(config.Content.LearningPaths, os.ModePerm); err != nil {
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

	// Create auxiliar structure for easy access to learning path books lpBooksData["apis"] = [{book1}, {book2}, ...]
	// Books are ordered by order and weight
	lpBooksData := map[sources.LearningPathRef][]sources.Book{}
	for _, b := range sources.Books {
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
	for _, b := range config.Badges {
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

	// Create template rendering data
	var data = struct {
		LpData              map[string]interface{}
		LpBooksData         []sources.Book
		BooksData           map[string]sources.Book
		BadgesData          map[string]interface{}
		BookCovers          string
		LearningPathsFolder string
		BooksIndex          string
		CurrentLearningPath sources.LearningPath
	}{
		LpData:              lpData,
		BooksData:           booksData,
		BookCovers:          config.Sources.BookCovers,
		BadgesData:          badgesData,
		LearningPathsFolder: config.Content.LearningPaths,
		BooksIndex:          config.Content.BookIndex,
	}

	// Load templates and functions
	templates, err := template.New("base").Funcs(sprig.TxtFuncMap()).ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	// Render content
	file := "stdout" // if debug mode spit to stdout

	if slices.Contains(contents, "book-index") {
		log.Printf("rendering book index in %s", config.Content.BookIndex)
		if !debug {
			file = config.Content.BookIndex
		}

		if err = content.Render(templates, "book-index.md.tpl", file, data); err != nil {
			log.Fatalln(err)
		}
	}

	if slices.Contains(contents, "readme") {
		log.Printf("rendering readme in %s", config.Content.Readme)
		if !debug {
			file = config.Content.Readme
		}

		if err = content.Render(templates, "readme.md.tpl", file, data); err != nil {
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
					file = filepath.Join(config.Content.LearningPaths, fmt.Sprintf("%s.md", lp.Ref))
				}
				log.Printf("rendering learning-path %s in %s", lp.Ref, file)

				if err = content.Render(templates, "learning-path.md.tpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	log.Println("done.")
}
