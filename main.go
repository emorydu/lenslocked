package main

import (
	"fmt"
	"net/http"

	"github.com/emorydu/lenslocked/controllers"
	"github.com/emorydu/lenslocked/models"
	"github.com/emorydu/lenslocked/templates"
	"github.com/emorydu/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(
		views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	db, err := models.Open(models.DefaultPostgresConfig())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usersC := controllers.Users{
		UserService: &models.UserService{DB: db},
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}
