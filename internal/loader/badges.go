package loader

import (
	"log"
	"os"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
	"gopkg.in/yaml.v3"
)

// Badges["excellent"] = top
var Badges map[string]any = make(map[string]any)

func loadBadges(basepath string) error {
	log.Printf("load badges from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("badges files %s\n", files)

	// Load the content of the files and populate the var
	var content []models.BadgeCategory
	for _, f := range files {
		content, err = loadBadgesFile(f)
		if err != nil {
			return err
		}

		for _, badge := range content {
			for _, b := range badge.Badges {
				Badges[string(b.Ref)] = b.Icon
				stats.SetTotalBadges(len(Badges))
			}
		}
	}

	return nil
}

func loadBadgesFile(path string) ([]models.BadgeCategory, error) {
	var content []models.BadgeCategory

	// Read the file
	raw, err := os.ReadFile(path)
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
