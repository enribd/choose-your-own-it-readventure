package models

type LearningPath struct {
	Name      string            `yaml:"name"`
	Ref       LearningPathRef   `yaml:"ref"`
	Status    string            `yaml:"status"`
	Desc      string            `yaml:"desc"`
	Summary   string            `yaml:"summary"`
	Tabs      []LearningPathTab `yaml:"tabs"`
	Related   []LearningPathRef `yaml:"related,omitempty"`
	Suggested []LearningPathRef `yaml:"suggested,omitempty"`
	Tags      []TagRef             `yaml:"tags,omitempty"`
	Logo      Logo              `yaml:"logo,omitempty"`
}

type LearningPathTab struct {
	Ref   LearningPathTabRef `yaml:"ref"`
  TabData LearningPathTabData `yaml:"data"`
}

type LearningPathTabData struct {
	Name  string `yaml:"name"`
	Icon  string `yaml:"icon,omitempty"`
	Desc  string `yaml:"desc,omitempty"`
	Order int    `yaml:"order,omitempty"`
}

type Logo struct {
	Source string `yaml:"source"`
	Height string `yaml:"height,omitempty"`
	Width  string `yaml:"width,omitempty"`
}

type LearningPathRef string
type LearningPathTabRef string
