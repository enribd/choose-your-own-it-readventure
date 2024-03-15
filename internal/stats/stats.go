package stats

type Stats struct {
	TotalBooks                int            // Books shown in content
	TotalSkippedBooks         int            // Books marked as draft
	TotalAuthors              int            // Authors shown in content
	TotalLearningPaths        int            // Learning paths shown in content
	TotalSkippedLearningPaths int            // Learning paths marked as coming soon
	TotalLearningPathBooks    map[string]int // Number of books in learning path
	TotalBooksRead            int            // Number of books read
}

// Global stats data
var Data Stats

// Initialize stats data
func New() {
	Data.TotalLearningPathBooks = make(map[string]int)
}

func SetTotalAuthors(total int) {
	Data.TotalAuthors = total
}

func SetTotalBooks(total int) {
	Data.TotalBooks = total
}

func IncSkippedBook() {
	Data.TotalSkippedBooks++
}

func SetTotalLearningPaths(total int) {
	Data.TotalLearningPaths = total
}

func IncSkippedLearningPath() {
	Data.TotalSkippedLearningPaths++
}

func SetTotalLearningPathBooks(lpRef string, total int) {
	Data.TotalLearningPathBooks[lpRef] = total
}

func SetTotalBooksRead(total int) {
	Data.TotalBooksRead = total
}
