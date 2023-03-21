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
				stats.SetTotalLearningPaths(len(LearningPaths))
			}
		}
	}

	// Avoid learning paths having empty related or suggested learning paths
	purgeEmtpyRelatedLearningPaths()
	purgeEmtpySuggestedLearningPaths()

	// Build auxiliar template structure
	for lpRef, lp := range LearningPaths {
		LearningPathsTmpl[string(lpRef)] = lp
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

// Avoid learning paths having empty related learning paths
func purgeEmtpyRelatedLearningPaths() {
	for _, lp := range LearningPaths {
		// Remove empty related paths
		var notEmtpyLPs []models.LearningPathRef
		for _, relatedRef := range lp.Related {
			if len(LearningPathBooks[relatedRef]) > 0 {
				notEmtpyLPs = append(notEmtpyLPs, relatedRef)
			} /* else {
							log.Printf("'%s' is an empty or a coming soon learning path, removed from '%s' related paths", relatedRef, lp.Ref)
			      } */
		}
		lp.Related = notEmtpyLPs
		LearningPaths[string(lp.Ref)] = lp
	}
}

// Avoid learning paths having empty suggested learning paths
func purgeEmtpySuggestedLearningPaths() {
	for _, lp := range LearningPaths {
		var notEmtpyLPs []models.LearningPathRef
		for _, suggestedRef := range lp.Suggested {
			// If the suggested exists in the active lps map add it to the new suggested list
			if _, ok := LearningPaths[string(suggestedRef)]; ok {
				notEmtpyLPs = append(notEmtpyLPs, suggestedRef)
			} /* else {
							log.Printf("'%s' is an empty or a coming soon learning path, removed from '%s' suggested paths", suggestedRef, lp.Ref)
			} */
		}
		lp.Suggested = notEmtpyLPs
		LearningPaths[string(lp.Ref)] = lp
	}
}
