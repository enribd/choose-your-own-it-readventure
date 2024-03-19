package models

// Book data
type Book struct {
	Cover         string             `yaml:"cover"`
	Title         string             `yaml:"title"`
	Subtitle      string             `yaml:"subtitle"`
	Draft         bool               `yaml:"draft"`
	Url           string             `yaml:"url"`
	Authors       []string           `yaml:"authors"`
	Release       string             `yaml:"release"`
	Pages         string             `yaml:"pages"`
	Desc          string             `yaml:"desc"`
	LearningPaths []BookLearningPath `yaml:"learning_paths"`
	BadgesRefs    []BadgeRef         `yaml:"badges"`
	Tags          []TagRef           `yaml:"tags,omitempty"`
}

// Book location in learning path
type BookLearningPath struct {
	// Reference to the learning path
	LearningPathRef LearningPathRef `yaml:"lp_ref"`

	// Tab of the learning path
	TabRef LearningPathTabRef `yaml:"tab_ref"`

	// Index of the book in the learning path
	Order int `yaml:"order"`

	// Resolves confict when more than one book has the same order in the same learning path tab
	Weight int `yaml:"weight"`
}

// Deduplicate slice that contains structs by a property
func DeduplicateBookLearningPaths(lps []BookLearningPath) []BookLearningPath {
	seen := make(map[LearningPathRef]bool)
	deduplicated := []BookLearningPath{}

	for _, lp := range lps {
		if !seen[lp.LearningPathRef] {
			seen[lp.LearningPathRef] = true
			deduplicated = append(deduplicated, lp)
		}
	}

	return deduplicated
}

