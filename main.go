package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	LogLevel      string         `yaml:"log_level"`
	Layout        Layout         `yaml:"layout"`
	LearningPaths []LearningPath `yaml:"learning_paths"`
	Badges        []Badge        `yaml:"badges"`
	Books         []Book         `yaml:"books"`
}

type Layout struct {
	Readme        string `yaml:"readme"`
	BookIndex     string `yaml:"book_index"`
	LearningPaths string `yaml:"learning_paths"`
	BookCovers    string `yaml:"book_covers"`
}

type LearningPath struct {
	Name         string            `yaml:"name"`
	Ref          string            `yaml:"reference"`
	Status       string            `yaml:"status"`
	Desc         string            `yaml:"description"`
	RelatedPaths []LearningPathRef `yaml:"related_paths,omitempty"`
}

type LearningPathRef string

type Badge struct {
	Category   Category    `yaml:"category"`
	BadgeIcons []BadgeIcon `yaml:"icons"`
}

type Category string

type BadgeIcon struct {
	Name string `yaml:"name"`
	Code string `yaml:"code"`
}

type Book struct {
	Cover             string            `yaml:"cover"`
	Title             string            `yaml:"title"`
	Subtitle          string            `yaml:"subtitle"`
	Order             string            `yaml:"order"`
	Draft             bool              `yaml:"draft"`
	Url               string            `yaml:"url"`
	Authors           []string          `yaml:"authors"`
	Release           string            `yaml:"release"`
	Pages             string            `yaml:"pages"`
	Description       string            `yaml:"description"`
	LearningPathsRefs []LearningPathRef `yaml:"learning_paths"`
	BookBadges        []BookBadge       `yaml:"badges"`
}

type BookBadge struct {
	Category Category `yaml:"category"`
	Value    string   `yaml:"value"`
}

var debug bool

func main() {
	// TODO segregate files, in layout indicate where are the yamls to load and do it
	// Read the file
	raw, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config

	// Unmarshal the YAML raw content into the struct
	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	if config.LogLevel == "debug" {
		debug = true
	}

	if debug {
		log.Printf("%v\n", config)
	}

	// Create auxiliar structure for easy access to learning paths lpData["apis"].Desc
	lpData := map[string]LearningPath{}
	for _, lp := range config.LearningPaths {
		lpData[lp.Ref] = lp
	}

	// Create auxiliar structure for easy access to books booksData["Building Microservices"].Desc
	booksData := map[string]Book{}
	for _, b := range config.Books {
		booksData[b.Title] = b
	}

	// Create auxiliar structure for easy access to learning path books lpBooksData["apis"] = [{book1}, {book2}, ...]
	lpBooksData := map[LearningPathRef][]Book{}
	for _, b := range config.Books {
		for _, r := range b.LearningPathsRefs {
			lpBooksData[r] = append(lpBooksData[r], b)
		}
	}

	// Create auxiliar structure for easy access to badges badgesData["rating"]["excellent"] = top
	badgesData := map[Category]map[string]string{}
	for _, b := range config.Badges {
		badgesData[b.Category] = map[string]string{}
		for _, i := range b.BadgeIcons {
			badgesData[b.Category][i.Name] = i.Code
		}
	}

	if debug {
		log.Printf("loaded learning paths: %v\n", lpData)
		log.Printf("loaded learning paths books: %v\n", lpBooksData)
		log.Printf("loaded books: %v\n", booksData)
		log.Printf("loaded badges: %v\n", badgesData)
	}

	// Create template rendering data
	var data = struct {
		LpData      map[string]LearningPath
		LpBooksData map[LearningPathRef][]Book
		BooksData   map[string]Book
		BadgesData  map[Category]map[string]string
	}{
		LpData:      lpData,
		LpBooksData: lpBooksData,
		BooksData:   booksData,
		BadgesData:  badgesData,
	}

	// Load templates and functions
	templates, err := template.New("base").Funcs(sprig.FuncMap()).ParseGlob("templates/*")
	if err != nil {
	  log.Fatalln(err)
	}

  if err = render(templates, "book-index.md.tpl", config.Layout.BookIndex, data); err != nil {
	  log.Fatalln(err)
  }

  if err = render(templates, "readme.md.tpl", config.Layout.Readme, data); err != nil{
	  log.Fatalln(err)
  }

  for _, lp := range config.LearningPaths {
    file := filepath.Join(config.Layout.LearningPaths, lp.Ref + ".md")
    if err = render(templates, "learning-path.md.tpl", file, data); err != nil {
	    log.Fatalln(err)
    }
  }
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
	// Create destination file
	var file *os.File
  var err error

	if debug {
	  file = os.Stdout
	} else {
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
