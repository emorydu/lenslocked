package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, tmpl string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", tmpl+".gohtml")
	t, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.
			StatusInternalServerError)

		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.
			StatusInternalServerError)

		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "home")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "contact")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "faq")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}
