{{- $lpBooksData := .LpBooksData -}}
{{- $lpData := .LpData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}
{{- $lp := .CurrentLearningPath -}}
# {{ $lp.Name }} Learning Path

{{ $lp.Desc }}
| Order | Badges | Cover | Info | Description |
| --- | --- | --- | --- | --- |
{{- range $lpBooksData }}
{{/* Build book badges section */}}
{{- $badges := list -}}
{{- range .Badges -}}
{{- $b := get $badgesData .Value | printf ":%s:" -}}
{{- $badges = append $badges $b -}}
{{- end -}}
{{- /* end Build book badges section */ -}}
| **{{ .Order }}** | {{ $badges | join " " }} | ![img]({{ if (.Cover | hasPrefix "http") }}{{ .Cover }}{{ else }}{{$covers}}/{{ .Cover }}{{end}}) | [**{{ .Title }}**]({{ .Url }}) <br> *{{ .Authors | join ", " }}* <br> *Published in {{ .Release }}* <br> *{{ .Pages }} pages* | {{ .Desc }} |
{{ end }}

{{- with $lp.Related }}
The following paths are opened to you now, choose wisely:

{{ range $lp.Related -}}
{{ $relPath := . | toString | get $lpData -}}
{{ if (eq $relPath.Status "coming-soon") -}}
- :{{ get $badgesData $relPath.Status }}: {{ $relPath.Name }}: {{ $relPath.Summary }}
{{- else -}}
{{- end -}}
- [{{ $relPath.Name }}]({{ $lpFolders }}/{{ $relPath.Ref }}.md): {{ $relPath.Summary }}
{{- end -}}
{{- end }}

{{- with $lp.Suggested }}
Want to change the subject? Here are some suggestions about other paths you can explore:

{{ range $lp.Suggested -}}
{{ $sugPath := . | toString | get $lpData -}}
{{ if (ne $sugPath.Status "coming-soon") -}}
- [{{ $sugPath.Name }}]({{ $lpFolders }}/{{ $sugPath.Ref }}.md): {{ $sugPath.Summary }}
{{- end -}}
{{- end -}}
{{- end }}
