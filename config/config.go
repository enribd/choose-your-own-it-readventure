package config

import (
	"io/ioutil"

	"github.com/enribd/choose-your-own-it-readventure/internal/content"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SiteUrl string                      `yaml:"site_url"`
	Sources Sources                     `yaml:"sources"`
	Content map[content.Provider]Format `yaml:"content"`
}

// Location of data used to generate content
type Sources struct {
	BookCovers    string `yaml:"books_covers"`
	BookData      string `yaml:"books_data"`
	LearningPaths string `yaml:"learning_paths_data"`
	BadgesData    string `yaml:"badges_data"`
}

// Destination for generated content
type Format struct {
	Index         string `yaml:"index"`
	BookIndex     string `yaml:"book_index"`
	AuthorIndex   string `yaml:"author_index"`
	LearningPaths string `yaml:"learning_paths"`
	Badges        string `yaml:"badges"`
	About         string `yaml:"about"`
	Mentions      string `yaml:"mentions"`
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
