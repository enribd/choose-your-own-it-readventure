package loader

import (
	"io/ioutil"
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"gopkg.in/yaml.v3"
)

func loadBadges(basepath string) error {
	log.Printf("load badges from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("badges files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []models.Badge
	for _, f := range files {
		content, err = loadBadgesFile(f)
		if err != nil {
			return err
		}

		for _, badge := range content {
			for _, i := range badge.BadgeIcons {
				Badges[i.Name] = i.Code
			}
		}
	}

	return nil
}

func loadBadgesFile(path string) ([]models.Badge, error) {
	var content []models.Badge

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
