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
		for i, lpRef := range book.LearningPathsRefs {
			if _, ok := LearningPaths[string(lpRef)]; !ok {
				book.LearningPathsRefs = append(book.LearningPathsRefs[:i], book.LearningPathsRefs[i+1:]...)
				// log.Printf("'%s' is an empty or a coming soon learning path, removed learning path reference from '%s' book", lpRef, book.Title)

				// Remove the book from the authors
				for _, name := range book.Authors {
					for h, b := range Authors[name] {
						if b.Title == book.Title {
							Authors[name] = append(Authors[name][:h], Authors[name][h+1:]...)
							// log.Printf("'%s' book is removed, delete it from author '%s' too", b.Title, name)
						}
					}
					// If an author has 0 books then remove it
					if len(Authors[name]) == 0 {
						delete(Authors, name)
						stats.SetTotalAuthors(len(Authors))
					}
				}
			}
		}

		// if the book doesn't have any learning path remove it
		if len(book.LearningPathsRefs) == 0 {
			delete(Books, book.Title)
			stats.SetTotalBooks(len(Books))
		}
	}
}
