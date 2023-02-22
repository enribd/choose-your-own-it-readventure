{{- $lpData := .LpData -}}
{{- $books := .BooksData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}

[//]: # (Auto generated file from templates)

# Book Index

| Cover | Info | Learning Paths | Badges |
| --- | --- | --- | --- |
{{- range $books -}}
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
| ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{ $covers | trimPrefix "." }}/{{ .Cover }}{{end}}) | [**{{ .Title }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* | {{if gt ($paths | len) 0 }}<ul>{{ $paths | join "" }}</ul>{{ end }} | {{ $badges | join " " }} |
{{- end }}

[**⬆ top**](#book-index)
