package content

import (
	"io/ioutil"
	"path/filepath"

	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v2"
)

type LearningPath struct {
	Name      string            `yaml:"name"`
	Ref       LearningPathRef   `yaml:"ref"`
	Status    string            `yaml:"status"`
	Desc      string            `yaml:"desc"`
	Summary   string            `yaml:"summary"`
	Related   []LearningPathRef `yaml:"related,omitempty"`
	Suggested []LearningPathRef `yaml:"suggested,omitempty"`
}

type LearningPathRef string

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
	Desc              string            `yaml:"desc"`
	LearningPathsRefs []LearningPathRef `yaml:"learning_paths"`
	BadgesRefs        []BadgeRef        `yaml:"badges"`
}

type BadgeRef string

var LearningPaths *map[string]LearningPath
var Books *map[string]Book

func LoadBooks(basepath string) error {
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the Books var
	var content interface{}
	for _, f := range files {
		content, err = loadFile(f)
		if err != nil {
			return err
		}

		c := content.(*map[string]Book)
		maps.Copy(Books, c)
	}

	return nil
}

func LoadLearningPaths(basepath string) error {
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the Books var
	var content interface{}
	for _, f := range files {
		content, err = loadFile(f)
		if err != nil {
			return err
		}

		c := content.(*map[string]LearningPath)
		maps.Copy(LearningPaths, c)
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

func loadFile(path string) (interface{}, error) {
	var content interface{}

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
