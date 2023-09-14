package main

import (
	"fmt"
	"github.com/emorydu/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, tmpl string) {
	t, err := views.Parse(tmpl)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.
			StatusInternalServerError)

		return
	}
	t.Execute(w, nil)
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
