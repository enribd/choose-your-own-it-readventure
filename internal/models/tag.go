package models

type TagRef string

type Tag struct {
	Ref  TagRef `yaml:"ref"`
	Name string `yaml:"name,omitempty"`
	Url  string `yaml:"url,omitempty"`
}
