package loader

import (
	"path/filepath"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
)

// LearningPaths["ref"] = LearningPath{}
var LearningPaths map[string]models.LearningPath = make(map[string]models.LearningPath)

// This structure is the same as LearningPaths but type agnostic, used in
// template data because it only allows map[string]interface{} types
var LearningPathsTmpl map[string]interface{} = make(map[string]interface{})

// Books["title"] = Book{}
var Books map[string]models.Book = make(map[string]models.Book)

// Authors["name"] = []Book{}
var Authors map[string][]models.Book = make(map[string][]models.Book)

// LearningPathBooks["ref"] = []Book{}
var LearningPathBooks map[models.LearningPathRef][]models.Book = make(map[models.LearningPathRef][]models.Book)

// Badges["excellent"] = top
var Badges map[string]interface{} = make(map[string]interface{})

func Load(booksPath, lpsPath, badgesPath string) error {
	err := loadBooks(booksPath)
	if err != nil {
		return err
	}

	loadAuthors()

	// Always load learning paths after books because learning paths without books are skipped.
	err = loadLearningPaths(lpsPath)
	if err != nil {
		return err
	}

	// Remove empty learning paths from books
	purgeEmtpyLearningPathRefsFromBooks()

	err = loadBadges(badgesPath)
	if err != nil {
		return err
	}

	return nil
}

func getFiles(basepath string) ([]string, error) {
	// Get all yaml files
	files, err := filepath.Glob(filepath.Join(basepath, "*.yaml"))
	if err != nil {
		return nil, err
	}
	return files, err
}
