package config

import (
	"io/ioutil"

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
	AuthorIndex   string `yaml:"author_index"`
	LearningPaths string `yaml:"learning_paths"`
}

// Icons by category
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

var Cfg Config

func Load() error {
	// Read the file
	raw, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	// Unmarshal the YAML raw content into the struct
	err = yaml.Unmarshal(raw, &Cfg)
	if err != nil {
		return err
	}

	return nil
}
