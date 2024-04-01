package loader

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load(booksPath, lpsPath, lpsTabsPath, badgesPath, tagsPath string) error {
	// Load before books because books badges are checked on book load
	err := loadBadges(badgesPath)
	if err != nil {
		return err
	}

	// Load before books because books tags are checked on book load
	err = loadTags(tagsPath)
	if err != nil {
		return err
	}

	err = loadBooks(booksPath)
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

func loadFile[T any](path string) ([]T, error) {
	var content []T

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
