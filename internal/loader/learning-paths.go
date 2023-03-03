package loader

import (
	"io/ioutil"
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
	"gopkg.in/yaml.v3"
)

func loadLearningPaths(basepath string) error {
	log.Printf("load learning paths from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("learning path files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []models.LearningPath
	for _, f := range files {
		content, err = loadLearningPathsFile(f)
		if err != nil {
			return err
		}

		for _, lp := range content {
			// skip if lp has no books or if status is coming soon
			if lp.Status == "coming-soon" || len(LearningPathBooks[lp.Ref]) == 0 {
				stats.IncSkippedLearningPath()
			} else {
				LearningPaths[string(lp.Ref)] = lp
				LearningPathsTmpl[string(lp.Ref)] = lp
				stats.SetTotalLearningPaths(len(LearningPaths))
			}
		}
	}

	return nil
}

func loadLearningPathsFile(path string) ([]models.LearningPath, error) {
	var content []models.LearningPath

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
