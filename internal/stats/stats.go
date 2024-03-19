package stats

// Global stats data
var Data Stats

type Stats struct {
	TotalBooks                int            // Books shown in content
	TotalSkippedBooks         int            // Books marked as draft
	TotalAuthors              int            // Authors shown in content
	TotalTags                 int            // Tags loaded from file
	TotalBadges               int            // Badges loaded from file
	TotalLearningPaths        int            // Learning paths shown in content
	TotalSkippedLearningPaths int            // Learning paths marked as coming soon
	TotalLearningPathBooks    map[string]int // Number of books in learning path
	TotalBooksRead            int            // Number of books read
}

// Initialize stats data
func New() {
	Data.TotalLearningPathBooks = make(map[string]int)
}

func SetTotalAuthors(total int) {
	Data.TotalAuthors = total
}

func SetTotalTags(total int) {
	Data.TotalTags = total
}

func SetTotalBadges(total int) {
	Data.TotalBadges = total
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
