package main

import (
	"fmt"
	"github.com/emorydu/lenslocked/controllers"
	"github.com/emorydu/lenslocked/models"
	"github.com/emorydu/lenslocked/templates"
	"github.com/emorydu/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"net/http"
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
		UserService:    &models.UserService{DB: db},
		SessionService: &models.SessionService{DB: db},
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/users", usersC.Create)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on: 3000...")

	csrfKey := "gFvi45R4fy5xNBlnEnZtQbfAVCYEIAUX"
	csrfMiddleware := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix the before deploying.
		csrf.Secure(false))
	http.ListenAndServe(":3000", csrfMiddleware(r))
	//http.ListenAndServe(":3000", TimerMiddleware(r.ServeHTTP))
}

//func TimerMiddleware(f http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//		f(w, r)
//		fmt.Println(r.Method, r.URL.Path, "Request Time", time.Since(start))
//	}
//}
