package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,
		"<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"orangeduxiaocheng@gmail.com\">orangeduxiaocheng@gmail.com</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/path", pathHandler)
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", nil)
}
