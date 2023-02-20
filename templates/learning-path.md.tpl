{{- $lpBooksData := .LpBooksData -}}
{{- $lpData := .LpData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}
{{- $lp := .CurrentLearningPath -}}

[//]: # (Auto generated file from templates)

# {{ $lp.Name }} Learning Path

{{ $lp.Desc }}
| Order | Badges | Cover | Info | Description |
| --- | --- | --- | --- | --- |
{{- range $lpBooksData }}
{{/* Build book badges section */}}
{{- $badges := list -}}
{{- range .BadgesRefs -}}
{{- $b := . | toString | get $badgesData | printf ":%s:" -}}
{{- $badges = append $badges $b -}}
{{- end -}}
{{- /* end Build book badges section */ -}}
| **{{ .Order }}** | {{ $badges | join " " }} | ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{$covers}}/{{ .Cover }}{{end}}) | [**{{ .Title }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* | {{ .Desc }} |
{{ end }}

{{- with $lp.Related }}
The following paths are opened to you now, choose wisely:

{{ range $lp.Related -}}
{{ $relPath := . | toString | get $lpData -}}
{{ $relPathIcon :=  get $badgesData $relPath.Status | printf ":%s:" -}}
{{ if (eq $relPath.Status "coming-soon") -}}
- :{{ get $badgesData $relPath.Status }}: {{ $relPath.Name }}: {{ $relPath.Summary }}
{{- else -}}
- [{{ $relPath.Name }} {{ $relPathIcon }}]({{ $lpFolders }}/{{ $relPath.Ref }}.md): {{ $relPath.Summary }}
{{- end -}}
{{- end -}}
{{- end }}

{{- with $lp.Suggested }}

Want to change the subject? Here are some suggestions about other paths you can explore:
{{ range $lp.Suggested -}}
{{ $sugPath := . | toString | get $lpData -}}
{{ $sugPathIcon :=  get $badgesData $sugPath.Status | printf ":%s:" -}}
{{ if (ne $sugPath.Status "coming-soon") }}
- [{{ $sugPath.Name }} {{ $sugPathIcon }}]({{ $lpFolders }}/{{ $sugPath.Ref }}.md): {{ $sugPath.Summary }}
{{- end -}}
{{- end -}}
{{- end }}

[**⬆ top**](#{{ $lp.Name | lower | replace " " "-" }}-learning-path)
