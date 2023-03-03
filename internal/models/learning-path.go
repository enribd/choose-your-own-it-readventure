package models

type LearningPathRef string

type Tag string

type Logo struct {
	Source string `yaml:"source"`
	Height string `yaml:"height,omitempty"`
	Width  string `yaml:"width,omitempty"`
}

type LearningPath struct {
	Name      string            `yaml:"name"`
	Ref       LearningPathRef   `yaml:"ref"`
	Status    string            `yaml:"status"`
	Desc      string            `yaml:"desc"`
	Summary   string            `yaml:"summary"`
	Related   []LearningPathRef `yaml:"related,omitempty"`
	Suggested []LearningPathRef `yaml:"suggested,omitempty"`
	Tags      []Tag             `yaml:"tags,omitempty"`
	Logo      Logo              `yaml:"logo,omitempty"`
}
