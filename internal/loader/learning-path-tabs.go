package loader

import (
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
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
		content, err = loadFile[models.LearningPathTab](f)
		if err != nil {
			return err
		}

		for _, lpt := range content {
			LearningPathsTabs[lpt.Ref] = lpt.Data
		}
	}

	return nil
}
