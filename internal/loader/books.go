package loader

import (
	"io/ioutil"
	"log"
	"slices"
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
        if slices.Contains(book.BadgesRefs, "read") {
          stats.SetTotalBooksRead(stats.Data.TotalBooksRead + 1)
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

func loadAuthors() {
	// Extract authors data
	for _, book := range Books {
		for _, name := range book.Authors {
			Authors[name] = append(Authors[name], book)
			stats.SetTotalAuthors(len(Authors))
		}
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

// Remove empty learning paths from books
func purgeEmtpyLearningPathRefsFromBooks() {
	for _, book := range Books {
		// Build a new list without empty lp refs
		var notEmtpyLPRefs []models.LearningPathRef

		for _, lpRef := range book.LearningPathsRefs {
			if _, ok := LearningPaths[string(lpRef)]; ok {
				notEmtpyLPRefs = append(notEmtpyLPRefs, lpRef)
			} /* else {
					// log.Printf("'%s' is an empty or a coming soon learning path, removed learning path reference from '%s' book", lpRef, book.Title)
			  } */
		}
		book.LearningPathsRefs = notEmtpyLPRefs

		// if the book doesn't have any learning path remove it
		if len(book.LearningPathsRefs) == 0 {
			delete(Books, book.Title)
			stats.SetTotalBooks(len(Books))
		}
	}
}
