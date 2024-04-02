package loader

import (
	"log"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
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
		content, err = loadFile[models.Tag](f)
		if err != nil {
			return err
		}

		// Create auxiliar map to store tags
		for _, tag := range content {
			// Check for duplicated tags
			if _, ok := Tags[string(tag.Ref)]; ok {
				log.Fatalf("loader: duplicated tag definition detected: %s", tag.Ref)
			}

			Tags[string(tag.Ref)] = tag
			stats.SetTotalTags(len(Tags))
		}
	}

	return nil
}
