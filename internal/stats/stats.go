package stats

type Stats struct {
	TotalBooks          int
	TotalAuthors        int
	TotalLearningPaths  int
	BooksInLearningPath map[string]int
}

var Data Stats

func New(tb, ta, tlp int, bilp map[string]int) {
	Data.TotalBooks = tb
	Data.TotalAuthors = ta
	Data.TotalLearningPaths = tlp
	Data.BooksInLearningPath = bilp
}
