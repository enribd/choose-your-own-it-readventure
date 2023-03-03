package sources

type LearningPathRef string
type BadgeRef string
type Tag string

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

type Book struct {
	Cover             string            `yaml:"cover"`
	Title             string            `yaml:"title"`
	Subtitle          string            `yaml:"subtitle"`
	Order             int               `yaml:"order"`
	Weight            int               `yaml:"weight"`
	Draft             bool              `yaml:"draft"`
	Url               string            `yaml:"url"`
	Authors           []string          `yaml:"authors"`
	Release           string            `yaml:"release"`
	Pages             string            `yaml:"pages"`
	Desc              string            `yaml:"desc"`
	LearningPathsRefs []LearningPathRef `yaml:"learning_paths"`
	BadgesRefs        []BadgeRef        `yaml:"badges"`
}

type Logo struct {
	Source string `yaml:"source"`
	Height string `yaml:"height,omitempty"`
	Width  string `yaml:"width,omitempty"`
}
