% {{ .Title }}

![Logo]({{ .Logo }})

# {{ .Name }} ({{ .Alias }})

> {{ .Quote.Text }}
> <br>
> — {{ .Quote.Author }}

{{ .Desc }}

## Tech Stack
{{ range .TechStack }}
* {{ .Name }}
{{ end }}

## Licenses & Ceritification
|Name|Issuer|
|---|---|{{ range .Certs }}
|{{ .Name }}|{{ .Issuer }}
|{{ end }}
