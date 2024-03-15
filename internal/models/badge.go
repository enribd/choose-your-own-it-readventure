package models

// Icons by category
type Badge struct {
	Name   BadgeCategory `yaml:"name"`
	Category   BadgeCategory `yaml:"category"`
	BadgeIcons []BadgeIcon   `yaml:"icons"`
}

type BadgeCategory string

type BadgeIcon struct {
	Name string `yaml:"name"`
	Code string `yaml:"code"`
	Desc string `yaml:"desc"`
}
