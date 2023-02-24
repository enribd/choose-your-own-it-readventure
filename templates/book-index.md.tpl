{{- $lpData := .LpData -}}
{{- $books := .BooksData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}

[//]: # (Auto generated file from templates)

# Book Index

| Cover | Info | Learning Paths |
| :---: | :--- | :--- |
{{- range $books -}}
{{- if .Draft | not -}}
{{/* Build book learning paths section */}}
{{- $paths := list -}}
{{- range .LearningPathsRefs -}}
{{- $lp := get $lpData (. | toString) -}}
{{- if (empty $lp | not) -}}
{{- $name := printf "<li>[%s](%s/%s.md)</li>" $lp.Name $lpFolders . | trim -}}
{{- $paths = append $paths $name -}}
{{- end -}}
{{- end -}}
{{/* end Build book learning paths list */}}
{{/* Build book badges section */}}
{{- $badges := list -}}
{{- range .BadgesRefs -}}
{{- $b := get $badgesData (. | toString) -}}
{{- if (empty $b | not) -}}
{{- $badges = append $badges (printf ":%s:" $b) -}}
{{- end -}}
{{- end -}}
{{- /* end Build book badges section */ -}}
{{- $sub := empty .Subtitle | ternary "" (printf ": %s" .Subtitle) -}}
| ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{ $covers | trimPrefix "." }}/{{ .Cover }}{{end}}) | [**{{ .Title }}{{ $sub }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* <br> {{ $badges | join " " }} | {{if gt ($paths | len) 0 }}<ul>{{ $paths | join "" }}</ul>{{ end }} |
{{- end }}
{{- end }}

[**⬆ top**](#book-index)
