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
	// fmt.Fprintln(w, r.URL.Path)
	// fmt.Fprintln(w, r.URL.RawPath)
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// TODO: handle the page not found error
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "404 Not Found")
	}
	// if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	//  return
	// } else if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	//  return
	// }
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

func main() {
	var router http.HandlerFunc = pathHandler
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/path", pathHandler)
	// mux := Router{}
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", router)
}
