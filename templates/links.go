package templates

import (
	"html/template"
	"io"
)

type LinksPage struct {
	Links []string
}

func CreateLinksPage(links []string) LinksPage {
	return LinksPage{
		Links: links,
	}
}

func (p LinksPage) Render(w io.Writer) error {
	tmplFile := "templates/links.html.tmpl"
	tmpl := template.Must(template.ParseFiles(tmplFile))
	err := tmpl.Execute(w, p)
	return err
}
