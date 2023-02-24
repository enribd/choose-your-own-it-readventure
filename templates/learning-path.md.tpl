{{- $lpBooksData := .LpBooksData -}}
{{- $lpData := .LpData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder | trimPrefix "." -}}
{{- $lp := .CurrentLearningPath -}}

[//]: # (Auto generated file from templates)

# {{ $lp.Name }} Learning Path

{{ $lp.Desc }}
| Reading Order | Cover | Info | Description |
| :---: | :---: | :--- | :--- |
{{- range $lpBooksData }}
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
| **{{ .Order }}** | ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{ $covers | trimPrefix "." }}/{{ .Cover }}{{end}}) | [**{{ .Title }}{{ $sub }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* <br> {{ $badges | join " " }} | {{ .Desc }} |
{{- end }}

{{- with $lp.Related }}

The following paths are opened to you now, choose wisely:

{{ range $lp.Related -}}
{{ $relPath := get $lpData (. | toString) -}}
{{- if (empty $relPath | not) -}}
{{ $relPathIcon :=  get $badgesData $relPath.Status | printf ":%s:" -}}
{{ if (eq $relPath.Status "coming-soon") -}}
- :{{ get $badgesData $relPath.Status }}: {{ $relPath.Name }}: {{ $relPath.Summary }}
{{- else -}}
- [{{ $relPath.Name }} {{ $relPathIcon }}]({{ $lpFolders }}/{{ $relPath.Ref }}.md): {{ $relPath.Summary }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end }}

{{- with $lp.Suggested }}

Want to change the subject? Here are some suggestions about other paths you can explore:
{{ range $lp.Suggested -}}
{{ $sugPath := get $lpData (. | toString ) -}}
{{- if (empty $sugPath | not) -}}
{{ $sugPathIcon :=  printf "" -}}
{{- if (eq (get $badgesData $sugPath.Status) "stable" | not) -}}
{{ $sugPathIcon =  get $badgesData $sugPath.Status | printf " :%s:" -}}
{{- end -}}
{{ if (ne $sugPath.Status "coming-soon") }}
- [{{ $sugPath.Name }}{{ $sugPathIcon }}]({{ $lpFolders }}/{{ $sugPath.Ref }}.md): {{ $sugPath.Summary }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end }}

---
[**⬆ back to top**](#{{ $lp.Name | lower | replace " " "-" }}-learning-path)
