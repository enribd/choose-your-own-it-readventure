package loader

import (
	"log"
	"os"
	"slices"
	"sort"

	"github.com/enribd/choose-your-own-it-readventure/internal/models"
	"github.com/enribd/choose-your-own-it-readventure/internal/stats"
	"gopkg.in/yaml.v3"
)

// Books["title"] = Book{}
var Books map[string]models.Book = make(map[string]models.Book)

// Authors["name"] = []Book{}
var Authors map[string][]models.Book = make(map[string][]models.Book)

// LearningPathBooks["ref"] = []Book{}
var LearningPathBooks map[models.LearningPathRef][]models.Book = make(map[models.LearningPathRef][]models.Book)

// LearningPathTabBooks["lp_ref"]["tab_ref"] = []Book{}
var LearningPathTabBooks map[models.LearningPathRef]map[models.LearningPathTabRef][]models.Book = make(map[models.LearningPathRef]map[models.LearningPathTabRef][]models.Book)

// This structure is the same as LearningPaths but type agnostic, used in
// template data because it only allows map[string]any types
var LearningPathTabBooksTmpl map[string]any = make(map[string]any)

func loadBooks(basepath string) error {
	log.Printf("load books from %s\n", basepath)
	files, err := getFiles(basepath)
	if err != nil {
		return err
	}

	// Load the content of the files and populate the var
	var content []models.Book
	for _, f := range files {
		log.Printf("book files %s\n", files)
		content, err = loadBooksFile(f)
		if err != nil {
			return err
		}

		for _, book := range content {
			// If book is in draft mode skip it
			if book.Draft {
				stats.IncSkippedBook()
			} else {
				// check for duplicates
				seenTags := make(map[models.TagRef]bool)
				for _, t := range book.Tags {
					if _, ok := seenTags[t]; ok {
						log.Fatalf("loader: %s book has duplicated tags: %s", book.Title, t)
					}
					seenTags[t] = true
				}

				// Add book to content
				Books[book.Title] = book
				stats.SetTotalBooks(len(Books))
				if slices.Contains(book.BadgesRefs, "read") {
					stats.SetTotalBooksRead(stats.Data.TotalBooksRead + 1)
				}

				// Insert book in learning path tab
				for _, lp := range book.LearningPaths {
					// Warning: if the learning path doesn't have a tab with TabRef, the tab and the books it contains won't be shown
					// Create a temporal copy of the book to remove all other learning paths and simplify later the sorting operations
					bookCopy := book
					bookCopy.LearningPaths = []models.BookLearningPath{lp}
					// Initialize nested tb map if it doesn't exist
					tabsMap, ok := LearningPathTabBooks[lp.LearningPathRef]
					if !ok {
						tabsMap = make(map[models.LearningPathTabRef][]models.Book)
						LearningPathTabBooks[lp.LearningPathRef] = tabsMap
					}
					LearningPathTabBooks[lp.LearningPathRef][lp.TabRef] = append(LearningPathTabBooks[lp.LearningPathRef][lp.TabRef], bookCopy)
				}
			}
		}
	}

	checkDuplicatesLearningPathTabBooks()
	sortAndCountLearningPathTabBooks()

	// Build auxiliar template structure
	for lpRef, tabs := range LearningPathTabBooks {
		for tabRef, books := range tabs {
			tabMap, ok := LearningPathTabBooksTmpl[string(lpRef)]
			// Initialize nested tab map if it doesn't exist
			if !ok {
				tabMap = make(map[string]any)
				LearningPathTabBooksTmpl[string(lpRef)] = tabMap
			}
			tm := LearningPathTabBooksTmpl[string(lpRef)].(map[string]any)
			tm[string(tabRef)] = books
		}
	}

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

func checkDuplicatesLearningPathTabBooks() {
	for lpRef, tabs := range LearningPathTabBooks {
		for tabRef, books := range tabs {
			seenBooksInTab := make(map[string]bool)
			for _, b := range books {
				if _, ok := seenBooksInTab[b.Title]; ok {
					log.Fatalf("loader: %s book is duplicated in tab %s of learning path %s", b.Title, tabRef, lpRef)
				}
				seenBooksInTab[b.Title] = true
			}
		}
	}
}

// Count and sort books by order ascendant and heavier weight in each learning path tab
func sortAndCountLearningPathTabBooks() {
	for lpRef, tabs := range LearningPathTabBooks {
		seenBookInLp := make(map[string]bool)
		for tabRef, books := range tabs {
			// the list of books in the tab only have one learning path (the rest was purged when loading), so we can use the first element to get the learning path order
			sort.SliceStable(books, func(i, j int) bool {
				if books[i].LearningPaths[0].Order != books[j].LearningPaths[0].Order {
					return books[i].LearningPaths[0].Order < books[j].LearningPaths[0].Order
				}

				// If orders are equal then sort by weight
				return books[i].LearningPaths[0].Weight > books[j].LearningPaths[0].Weight
			})

			// Count book only if first seen
			for _, b := range books {
				if _, ok := seenBookInLp[b.Title]; !ok {
					// Count (previous learning path total plus the current tab books)
					newTotal := stats.Data.TotalLearningPathBooks[string(lpRef)] + len(LearningPathTabBooks[lpRef][tabRef])
					stats.SetTotalLearningPathBooks(string(lpRef), newTotal)
					seenBookInLp[b.Title] = true
				}
			}
		}
	}
}

// Remove empty learning paths from books
func purgeEmtpyLearningPathRefsFromBooks() {
	for _, book := range Books {
		// Build a new list without empty lp refs
		var newBookLearningPaths []models.BookLearningPath

		for _, lp := range book.LearningPaths {
			if _, ok := LearningPaths[string(lp.LearningPathRef)]; ok {
				newBookLearningPaths = append(newBookLearningPaths, lp)
			} /* else {
					// log.Printf("'%s' is an empty or a coming soon learning path, removed learning path reference from '%s' book", lpRef, book.Title)
			  } */
		}

		book.LearningPaths = newBookLearningPaths

		// if the book doesn't have any learning paths remove it
		if len(book.LearningPaths) == 0 {
			delete(Books, book.Title)
			stats.SetTotalBooks(len(Books))
		}
	}
}
