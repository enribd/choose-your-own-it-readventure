package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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
var content string
var config Config

func main() {
	// Parse flags
	flag.BoolVar(&debug, "debug", false, "Enable debug mode (default: false).")
	flag.BoolVar(&trace, "trace", false, "Enable trace mode (default: false).")
	flag.StringVar(&content, "content", "readme,book-index,learning-paths", "list of content to generate, accepts comma-separated values")
	flag.Parse()
	contents := strings.Split(content, ",")

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

	// Create auxiliar structure for easy access to learning paths lpData["apis"].Desc
	lpData := map[string]interface{}{}
	for _, lp := range lpContent {
		lpData[string(lp.Ref)] = lp
	}

	// Create auxiliar structure for easy access to books booksData["Building Microservices"].Desc
	booksData := map[string]Book{}
	for _, b := range booksContent {
		booksData[b.Title] = b
	}

	// Create auxiliar structure for easy access to learning path books lpBooksData["apis"] = [{book1}, {book2}, ...]
	lpBooksData := map[LearningPathRef][]Book{}
	for _, b := range booksContent {
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

	// Create template rendering data
	var data = struct {
		LpData              map[string]interface{}
		LpBooksData         []Book
		BooksData           map[string]Book
		BadgesData          map[string]interface{}
		BookCovers          string
		LearningPathsFolder string
		BooksIndex          string
		CurrentLearningPath LearningPath
	}{
		LpData:              lpData,
		BooksData:           booksData,
		BookCovers:          config.Sources.BookCovers,
		BadgesData:          badgesData,
		LearningPathsFolder: config.Sources.LearningPaths,
		BooksIndex:          config.Content.BookIndex,
	}

	// Load templates and functions
	templates, err := template.New("base").Funcs(sprig.FuncMap()).ParseGlob("templates/*")
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
		for _, lp := range lpContent {
			if lp.Status != "coming-soon" {
				data.CurrentLearningPath = lp
				data.LpBooksData = lpBooksData[lp.Ref]

				// TODO remove -test from file name
				file := filepath.Join(config.Sources.LearningPaths, fmt.Sprintf("%s-test.md", lp.Ref))
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
