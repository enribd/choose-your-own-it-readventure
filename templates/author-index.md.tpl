[//]: # (Auto generated file from templates)

# Author Index

| Name | Books |
| :---: | :---: | :---: |
{{ $authorsData := .AuthorsData -}}
{{- range $author, $booksData := $authorsData -}}
{{- $books := list -}}
{{- range $booksData -}}
{{- $books = append $books (printf "<li>[*%s*](%s)</li>" .Title .Url) -}}
{{- end -}}
| {{ $author }} | <ul>{{ $books | join " " }}</ul> |
{{ end }}

[**⬆ top**](#author-index)
