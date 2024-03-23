package models

// Learning path data
type LearningPath struct {
	Name      string            `yaml:"name"`
	Ref       LearningPathRef   `yaml:"ref"`
	Status    string            `yaml:"status"`
	Desc      string            `yaml:"desc"`
	Summary   string            `yaml:"summary"`
	Tabs      []LearningPathTab `yaml:"tabs"`
	Related   []LearningPathRef `yaml:"related,omitempty"`
	Suggested []LearningPathRef `yaml:"suggested,omitempty"`
	Tags      []TagRef          `yaml:"tags,omitempty"`
	Logo      Logo              `yaml:"logo,omitempty"`
}

// Learning path tab spec
type LearningPathTab struct {
	Ref  LearningPathTabRef  `yaml:"ref"`
	Data LearningPathTabData `yaml:"data"`
}

// Learning path tab data
type LearningPathTabData struct {
	Name  string `yaml:"name"`
	Icon  string `yaml:"icon,omitempty"`
	Desc  string `yaml:"desc,omitempty"`
	Order int    `yaml:"order,omitempty"`
}

// Learning path logo
type Logo struct {
	Source string `yaml:"source"`
	Height string `yaml:"height,omitempty"`
	Width  string `yaml:"width,omitempty"`
}

// Learning path reference code
type LearningPathRef string

// Learning path tab reference code
type LearningPathTabRef string
