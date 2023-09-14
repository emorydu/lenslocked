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

	t, err := views.Parse("home")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(t))

	t, err = views.Parse("contact")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(t))

	t, err = views.Parse("faq")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(t))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}
