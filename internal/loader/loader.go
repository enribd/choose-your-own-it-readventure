package loader

import (
	"path/filepath"
)

func Load(booksPath, lpsPath, lpsTabsPath, badgesPath, tagsPath string) error {
	err := loadBooks(booksPath)
	if err != nil {
		return err
	}

	loadAuthors()

	// Load before learning paths because tabs are used in learning paths
	err = loadLearningPathsTabs(lpsTabsPath)
	if err != nil {
		return err
	}

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

	err = loadTags(tagsPath)
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
