package models

// Tag reference code
type TagRef string

// Tag data
type Tag struct {
	Ref  TagRef `yaml:"ref"`
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}
