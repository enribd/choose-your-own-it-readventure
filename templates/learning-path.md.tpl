{{- $lpBooksData := .LpBooksData -}}
{{- $covers := .BookCovers -}}
{{- $badgesData := .BadgesData -}}
{{- $lpFolders := .LearningPathsFolder -}}
{{- $lp := .CurrentLearningPath -}}
# {{ $lp.Name }} Learning Path

{{ $lp.Desc }}
| Order | Badges | Cover | Info | Description |
| --- | --- | ---| ---| --- |
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

{{- $p := len $lp.RelatedPaths -}}
{{ if (gt $p 0) }}
The following paths are opened to you now, choose wisely:

  - [Microservice Architectures](#microservice-architectures): Study the pinnacle of distributed systems architectures, learn its tenets, and foremost, when and how to implement it.
  - [API Design](#api-design): APIs are one way services use to talk to each other, there are a lot of aspects involved: communication protocols (REST, gRPC, WebSocket, GraphQL, etc), interface definition, version management, testing, security, rate limiting, patterns, api gateways, and more.
  - [Event Driven Architectures (EDA)](#event-driven-architectures-(eda)): Asynchronous communication between services is possible using events. There is a lot to learn here, the main challenge is changing the way you think about information distribution.
  - [Serverless](#serverless): Also known as Function as a Service (FaaS). It's a cloud-native development model and a computing paradigm that allows you to define your applications as functions and events and run them without provisioning or managing servers.
{{- end }}
