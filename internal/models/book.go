package models

type BadgeRef string

type Book struct {
	Cover             string                               `yaml:"cover"`
	Title             string                               `yaml:"title"`
	Subtitle          string                               `yaml:"subtitle"`
	Draft             bool                                 `yaml:"draft"`
	Url               string                               `yaml:"url"`
	Authors           []string                             `yaml:"authors"`
	Release           string                               `yaml:"release"`
	Pages             string                               `yaml:"pages"`
	Desc              string                               `yaml:"desc"`
	BookLearningPaths map[LearningPathRef]BookLearningPath `yaml:"learning_paths"`
	BadgesRefs        []BadgeRef                           `yaml:"badges"`
}

type BookLearningPath struct {
	LPRef  TabRef `yaml:"lp_ref"`
	TabRef TabRef `yaml:"tab_ref"`
	Order  int    `yaml:"order"`
	Weight int    `yaml:"weight"`
}
