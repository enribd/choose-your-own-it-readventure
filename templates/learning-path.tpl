{{ range .LearningPaths }}
# {{ .Name }} LearningPath

{{ .Desc }}

| Order | Badges | Cover | Info | Description |
| --- | --- | ---| ---| --- |
| **{{ .Order }}** | :dog2: :orange_book: :arrows_counterclockwise: | ![img]({{ if (.Cover hasPrefix "http") }}{{ .Cover }}{{ else }}{{.Layout.BookCovers}}/{{ .Book.Cover }}{{end}}) | [**{{ .Title }}**](https://www.oreilly.com/library/view/designing-distributed-systems/9781491983638/) <br> *Brendan Burns* <br> *Published in 2018* <br> *162 pages*                                                        | desc |
{{ end }}
