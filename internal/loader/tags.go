package loader

import (
	"os"
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
	"gopkg.in/yaml.v3"
)

// Tags["kubernetes"] = []Tag{}
var Tags map[string]any = make(map[string]any)

func loadTags(basepath string) error {
	log.Printf("load tags from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("tags files %s\n", files)

	// Load the content of the files and populate the var
	var content []models.Tag
	for _, f := range files {
		content, err = loadTagsFile(f)
		if err != nil {
			return err
		}

		// Create auxiliar map to store tags
		for _, tag := range content {
			Tags[string(tag.Ref)] = tag
			stats.SetTotalTags(len(Tags))
		}
	}

	return nil
}

func loadTagsFile(path string) ([]models.Tag, error) {
	var content []models.Tag

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
