package controllers

import (
	"html/template"
	"net/http"

	"github.com/emorydu/lenslocked/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version? ",
			Answer:   "We offer a free trial for 30 days on any paid plans.",
		}, {
			Question: "What are your support hours? ",
			Answer:   "We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support? ",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@orangeduxiaocheng@gmail.com</a>`,
		},
		{
			Question: "Where is your office located ",
			Answer:   "Our entire team is remote!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
