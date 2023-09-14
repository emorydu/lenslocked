package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(tmpl string) (Template, error) {
	tmplPath := filepath.Join("templates", tmpl+".gohtml")
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTmpl: t}, nil
}

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.
			StatusInternalServerError)
		return
	}
}
