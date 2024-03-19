package loader

import (
	"os"
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"gopkg.in/yaml.v3"
)

// LearningPathsTabs["ref"] = LearningPathTab{}
var LearningPathsTabs map[models.LearningPathTabRef]models.LearningPathTabData = make(map[models.LearningPathTabRef]models.LearningPathTabData)

func loadLearningPathsTabs(basepath string) error {
	log.Printf("load learning paths tabs from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("learning path tabs files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []models.LearningPathTab
	for _, f := range files {
		content, err = loadLearningPathsTabsFile(f)
		if err != nil {
			return err
		}

		for _, lpt := range content {
			LearningPathsTabs[lpt.Ref] = lpt.Data
		}
	}

	return nil
}

func loadLearningPathsTabsFile(path string) ([]models.LearningPathTab, error) {
	var content []models.LearningPathTab

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
