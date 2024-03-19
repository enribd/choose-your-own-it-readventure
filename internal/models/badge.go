package models

// List of badges by category
type BadgeCategory struct {
	Name   BadgeCategoryRef `yaml:"category"`
	Badges []Badge          `yaml:"badges"`
}

// Badge data
type Badge struct {
	Ref  BadgeRef `yaml:"ref"`
	Icon string   `yaml:"icon"`
	Desc string   `yaml:"desc,omitempty"`
}

// Badge category reference code
type BadgeCategoryRef string

// Badge reference code
type BadgeRef string
