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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
	<li><b>Is there a free version? Yes!</b> We offer a free trial for 30 days on any
  paid plans.</li>
	<li><b>What are your support hours?</b> We have support staff answering emails 24/7, though response times may be a
	bit slower on weekends.</li>
	<li><b>How do I contact support?</b> Email us - <a href="mailto:support@lenslocked.com">support@orangeduxiaocheng@gmail.com</a></li>
</ul>
`)
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
