package models

// Icons by category
// type Badge struct {
// 	Category   BadgeCategory `yaml:"category"`
// 	BadgeIcons []BadgeIcon   `yaml:"icons"`
// }
//
// type BadgeCategory string
//
// type BadgeIcon struct {
// 	Name string `yaml:"name"`
// 	Code string `yaml:"code"`
// 	Desc string `yaml:"desc"`
// }

type BadgeCategory struct {
	Name   BadgeCategory `yaml:"category"`
	Badges []Badge       `yaml:"badges"`
}

type Badge struct {
	Ref  BadgeRef `yaml:"ref"`
	Icon string   `yaml:"icon"`
	Desc string   `yaml:"desc"`
}

type BadgeCategory string
type BadgeRef string
