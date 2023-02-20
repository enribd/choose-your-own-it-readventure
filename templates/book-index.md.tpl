{{- $lpData := .LpData -}}
{{- $books := .BooksData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}
# Book Index

| Cover | Info | Learning Paths | Badges |
| --- | --- | --- | --- |
{{- range $books -}}
{{/* Build book learning paths section */}}
{{- $paths := list -}}
{{- range .LearningPathsRefs -}}
{{- $p := . | toString | get $lpData -}}
{{- $name := print "<li>[" $p.Name "](" $lpFolders "/" $p.Ref ".md)</li>" | trim -}}
{{- $paths = append $paths $name -}}
{{- end -}}
{{/* end Build book learning paths list */}}
{{/* Build book badges section */}}
{{- $badges := list -}}
{{- range .Badges -}}
{{- $b := get $badgesData .Value | printf ":%s:" -}}
{{- $badges = append $badges $b -}}
{{- end -}}
{{- /* end Build book badges section */ -}}
| ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{$covers}}/{{ .Cover }}{{end}}) | [**{{ .Title }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* | <ul>{{ $paths | join "" }}</ul> | {{ $badges | join " " }} |
{{- end }}

[**⬆ top**](#book-index)
