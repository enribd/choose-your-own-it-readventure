package sources

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type LearningPath struct {
	Name      string            `yaml:"name"`
	Ref       LearningPathRef   `yaml:"ref"`
	Status    string            `yaml:"status"`
	Desc      string            `yaml:"desc"`
	Summary   string            `yaml:"summary"`
	Related   []LearningPathRef `yaml:"related,omitempty"`
	Suggested []LearningPathRef `yaml:"suggested,omitempty"`
	Tags      []Tag             `yaml:"tags,omitempty"`
	Logo      Logo              `yaml:"logo,omitempty"`
}

type LearningPathRef string
type Tag string

type Book struct {
	Cover             string            `yaml:"cover"`
	Title             string            `yaml:"title"`
	Subtitle          string            `yaml:"subtitle"`
	Order             int               `yaml:"order"`
	Weight            int               `yaml:"weight"`
	Draft             bool              `yaml:"draft"`
	Url               string            `yaml:"url"`
	Authors           []string          `yaml:"authors"`
	Release           string            `yaml:"release"`
	Pages             string            `yaml:"pages"`
	Desc              string            `yaml:"desc"`
	LearningPathsRefs []LearningPathRef `yaml:"learning_paths"`
	BadgesRefs        []BadgeRef        `yaml:"badges"`
}

type BadgeRef string

type Logo struct {
	Source string `yaml:"source"`
	Height string `yaml:"height,omitempty"`
	Width  string `yaml:"width,omitempty"`
}

var LearningPaths map[string]LearningPath = make(map[string]LearningPath)
var Books map[string]Book = make(map[string]Book)

func LoadBooks(basepath string) error {
	log.Printf("load books from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the Books var
	var content []Book
	for _, f := range files {
		log.Printf("book files %s\n", files)
		content, err = loadBooksFile(f)
		if err != nil {
			return err
		}

		for _, book := range content {
			Books[book.Title] = book
		}
	}

	return nil
}

func LoadLearningPaths(basepath string) error {
	log.Printf("load learning paths from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("learning path files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []LearningPath
	for _, f := range files {
		content, err = loadLearningPathsFile(f)
		if err != nil {
			return err
		}

		for _, lp := range content {
			LearningPaths[string(lp.Ref)] = lp
		}
	}

	return nil
}

func getFiles(basepath string) ([]string, error) {
	// Get all yaml files
	files, err := filepath.Glob(filepath.Join(basepath, "*.yaml"))
	if err != nil {
		return nil, err
	}
	return files, err
}

func loadBooksFile(path string) ([]Book, error) {
	var content []Book

	// Read the file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML raw content into the struct
	err = yaml.Unmarshal(raw, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func loadLearningPathsFile(path string) ([]LearningPath, error) {
	var content []LearningPath

	// Read the file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML raw content into the struct
	err = yaml.Unmarshal(raw, &content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
