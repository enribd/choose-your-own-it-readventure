package sources

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"github.com/enribd/choose-your-own-it-readventure/internal/stats"

	"gopkg.in/yaml.v3"
)

// LearningPaths["ref"] = LearningPath{}
var LearningPaths map[string]LearningPath = make(map[string]LearningPath)

// This structure is the same as LearningPaths but type agnostic, used in
// template data because it only allows map[string]interface{} types
var LearningPathsTmpl map[string]interface{} = make(map[string]interface{})

// Books["title"] = Book{}
var Books map[string]Book = make(map[string]Book)

// Authors["name"] = []Book{}
var Authors map[string][]Book = make(map[string][]Book)

// LearningPathBooks["ref"] = []Book{}
var LearningPathBooks map[LearningPathRef][]Book = make(map[LearningPathRef][]Book)

func Load(booksPath, lpsPath string) error {
	err := loadBooks(booksPath)
	if err != nil {
		return err
	}

	// Always load learning paths after books because learning paths without books are skipped.
	err = loadLearningPaths(lpsPath)
	if err != nil {
		return err
	}

	return nil
}

func loadBooks(basepath string) error {
	log.Printf("load books from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the Books var
	var content []Book
	for _, f := range files {
		log.Printf("book files %s\n", files)
		content, err = loadBooksFile(f)
		if err != nil {
			return err
		}

		for _, book := range content {
			// if book is in draft mode skip it
			if book.Draft {
				stats.IncSkippedBook()
			} else {
				// Add book to content
				Books[book.Title] = book
				stats.SetTotalBooks(len(Books))

				// Extract authors data
				for _, name := range book.Authors {
					Authors[name] = append(Authors[name], book)
					stats.SetTotalAuthors(len(Authors))
				}

				// Insert book in learning path
				for _, lpRef := range book.LearningPathsRefs {
					LearningPathBooks[lpRef] = append(LearningPathBooks[lpRef], book)
				}
			}
		}
	}

	sortAndCountLearningPathBooks()

	return nil
}

// Sort and count books in each learning path by order ascendant and heavier weight
func sortAndCountLearningPathBooks() {
	for lpRef, books := range LearningPathBooks {
		// Sort
		sort.SliceStable(books, func(i, j int) bool {
			if books[i].Order != books[j].Order {
				return books[i].Order < books[j].Order
			}

			// If orders are equal then sort by weight
			return books[i].Weight > books[j].Weight
		})

		// Count
		stats.SetTotalLearningPathBooks(string(lpRef), len(LearningPathBooks[lpRef]))
	}
}

func loadLearningPaths(basepath string) error {
	log.Printf("load learning paths from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}
	log.Printf("learning path files %s\n", files)

	// Load the content of the files and populate the Books var
	var content []LearningPath
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

func getFiles(basepath string) ([]string, error) {
	// Get all yaml files
	files, err := filepath.Glob(filepath.Join(basepath, "*.yaml"))
	if err != nil {
		return nil, err
	}
	return files, err
}

func loadBooksFile(path string) ([]Book, error) {
	var content []Book

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

func loadLearningPathsFile(path string) ([]LearningPath, error) {
	var content []LearningPath

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
