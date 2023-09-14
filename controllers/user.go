package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	//err := r.ParseForm()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//}
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	_, _ = fmt.Fprintf(w, "Email: %s\n", email)
	_, _ = fmt.Fprintf(w, "Password: %s\n", pwd)

}
