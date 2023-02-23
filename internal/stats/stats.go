package stats

import "github.com/enribd/choose-your-own-it-readventure/internal/sources"

type Stats struct {
	TotalBooks          int
	TotalAuthors        int
	TotalLearningPaths  int
	BooksInLearningPath map[sources.LearningPathRef]int
}

var Data Stats

func New(tb, ta, tlp int, bilp map[sources.LearningPathRef]int) {
	Data.TotalBooks = tb
	Data.TotalAuthors = ta
	Data.TotalLearningPaths = tlp
	Data.BooksInLearningPath = bilp
}
