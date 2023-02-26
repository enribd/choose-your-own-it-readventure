package stats

type Stats struct {
	TotalBooks                int
	TotalActiveBooks          int
	TotalSkippedBooks         int
	TotalAuthors              int
	TotalLearningPaths        int
	TotalActiveLearningPaths  int
	TotalSkippedLearningPaths int
	BooksInLearningPath       map[string]int
}

var Data Stats

func New(tb, tab, tsb, ta, tlp, talp, tslp int, bilp map[string]int) {
	Data.TotalBooks = tb
	Data.TotalActiveBooks = tab
	Data.TotalSkippedBooks = tsb
	Data.TotalAuthors = ta
	Data.TotalLearningPaths = tlp
	Data.TotalActiveLearningPaths = tlp - tslp
	Data.TotalSkippedLearningPaths = tslp
	Data.BooksInLearningPath = bilp
}
