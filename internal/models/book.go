package models

type BadgeRef string

type Book struct {
	Cover             string            `yaml:"cover"`
	Title             string            `yaml:"title"`
	Subtitle          string            `yaml:"subtitle"`
	Draft             bool              `yaml:"draft"`
	Url               string            `yaml:"url"`
	Authors           []string          `yaml:"authors"`
	Release           string            `yaml:"release"`
	Pages             string            `yaml:"pages"`
	Desc              string            `yaml:"desc"`
	LearningPaths     []LearningPaths   `yaml:"learning_paths"`
	BadgesRefs        []BadgeRef        `yaml:"badges"`
}

type LearningPaths struct {
	LearningPathsRef  LearningPathRef   `yaml:"ref"`
	Order             int               `yaml:"order"`
	Weight            int               `yaml:"weight"`
}
