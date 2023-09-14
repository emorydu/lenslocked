package main

import (
	"fmt"
	"net/http"

	"github.com/emorydu/lenslocked/controllers"
	"github.com/emorydu/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse("home"))))
	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse("contact"))))
	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse("faq"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}
