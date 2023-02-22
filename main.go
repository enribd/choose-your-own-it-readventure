package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/enribd/choose-your-own-it-readventure/internal/content"
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
	// Parse flags
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

	content.LoadBooks(config.Sources.BookData)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("books: %T\n", content.Books)

	content.LoadLearningPaths(config.Sources.LearningPaths)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Printf("learning paths: %T\n", content.LearningPaths)

	// Create auxiliar structure for easy access to learning paths lpData["apis"].Desc
	lpData := map[string]interface{}{}
	for _, lp := range content.LearningPaths {
		lpData[string(lp.Ref)] = lp
	}
	// log.Printf("LPDATA: %v", lpData["apis"])

	// Create auxiliar structure for easy access to books booksData["Building Microservices"].Desc
	booksData := map[string]content.Book{}
	for _, b := range content.Books {
		booksData[b.Title] = b
	}

	// Create auxiliar structure for easy access to learning path books lpBooksData["apis"] = [{book1}, {book2}, ...]
	lpBooksData := map[content.LearningPathRef][]content.Book{}
	for _, b := range content.Books {
		for _, r := range b.LearningPathsRefs {
			lpBooksData[r] = append(lpBooksData[r], b)
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

	// Create content dirs
	if err = os.MkdirAll(config.Content.LearningPaths, os.ModePerm); err != nil {
		log.Println(err)
		return
	}

	// Create template rendering data
	var data = struct {
		LpData              map[string]interface{}
		LpBooksData         []content.Book
		BooksData           map[string]content.Book
		BadgesData          map[string]interface{}
		BookCovers          string
		LearningPathsFolder string
		BooksIndex          string
		CurrentLearningPath content.LearningPath
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

	if slices.Contains(contents, "book-index") {
		log.Printf("rendering book index in %s", config.Content.BookIndex)
		if err = render(templates, "book-index.md.tpl", config.Content.BookIndex, data); err != nil {
			log.Fatalln(err)
		}
	}

	if slices.Contains(contents, "readme") {
		log.Printf("rendering readme in %s", config.Content.Readme)
		if err = render(templates, "readme.md.tpl", config.Content.Readme, data); err != nil {
			log.Fatalln(err)
		}
	}

	if slices.Contains(contents, "learning-paths") {
		for _, lp := range content.LearningPaths {
			if lp.Status != "coming-soon" {
				data.CurrentLearningPath = lp
				data.LpBooksData = lpBooksData[lp.Ref]

				file := filepath.Join(config.Content.LearningPaths, fmt.Sprintf("%s.md", lp.Ref))
				log.Printf("rendering learning-path %s in %s", lp.Ref, file)

				if err = render(templates, "learning-path.md.tpl", file, data); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	log.Println("done.")
}

/*
* Render templates with a given data and export them to files or stdout
* Params:
*   t: templates loaded
*   data: data used to fill the templates
*   templateName: template name to render
*   dest: destination file
 */
func render(t *template.Template, templateName, dest string, data interface{}) error {
	var file *os.File
	var err error

	if debug {
		file = os.Stdout
	} else {
		// Create destination file
		file, err = os.Create(dest)
		if err != nil {
			log.Fatalln("create file: ", err)
			return err
		}
	}

	// Render template
	err = t.ExecuteTemplate(file, templateName, data)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
