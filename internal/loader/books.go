package loader

import (
	"io/ioutil"
	"log"
	"sort"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
	"gopkg.in/yaml.v3"
)

func loadBooks(basepath string) error {
	log.Printf("load books from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the Books var
	var content []models.Book
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

func loadBooksFile(path string) ([]models.Book, error) {
	var content []models.Book

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
