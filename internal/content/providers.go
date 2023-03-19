package content

import "fmt"

const (
	Github Provider = "github"
	Mkdocs Provider = "mkdocs"
)

type Provider string

func NewProvider(s string) (Provider, error) {
	switch s {
	case "github":
		return Github, nil
	case "mkdocs":
		return Mkdocs, nil
	}
	return "", fmt.Errorf("unkown provider %s, available: %s, %s", s, Github, Mkdocs)
}

func (p Provider) String() string {
	switch p {
	case Github:
		return "github"
	case Mkdocs:
		return "mkdocs"
	}
	return "unknown"
}
