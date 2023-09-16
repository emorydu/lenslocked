package views

import (
	"fmt"
	"github.com/gorilla/csrf"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	t := template.New(patterns[0])
	t = t.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return `<!-- TODO: Implement the csrfField -->`
		},
	})

	t, err := t.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTmpl: t}, nil
}

//func Parse(tmpl string) (Template, error) {
//	tmplPath := filepath.Join("templates", tmpl+".gohtml")
//	t, err := template.ParseFiles(tmplPath)
//	if err != nil {
//		return Template{}, fmt.Errorf("parsing template: %w", err)
//	}
//
//	return Template{htmlTmpl: t}, nil
//}

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//htmlTmpl := t.htmlTmpl	// TODO: Bug (Multiple Web Request [csrfFiled overwrite])
	htmlTmpl, err := t.htmlTmpl.Clone()
	if err != nil {
		log.Printf("cloning teamplate: %v", err)
		http.Error(w, "There was an error rendering the page.", http.
			StatusInternalServerError)
	}
	htmlTmpl = htmlTmpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})
	err = htmlTmpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.
			StatusInternalServerError)
		return
	}
}
