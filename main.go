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
		"<h1>Contact Page</h1><p>To get in touch, email me at <a href=\""+
			"rangeduxiaocheng@gmail.com\">orangeduxiaocheng@gmail.com</a>.")
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

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, r.URL.Path)
// 	// fmt.Fprintln(w, r.URL.RawPath)
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// TODO: handle the page not found error
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "404 Not Found")
// 	}
// 	// if r.URL.Path == "/" {
// 	// 	homeHandler(w, r)
// 	//  return
// 	// } else if r.URL.Path == "/contact" {
// 	// 	contactHandler(w, r)
// 	//  return
// 	// }
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

// type Server struct{}

// func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {

	// http.Handler - interface with the ServeHTTP method
	// http.HandlerFunc - a function type that accepts same args as ServeHTTP
	// method. also implements http.Handler

	// http.Handle("/", http.Handler)

	// var router http.HandlerFunc = pathHandler
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.HandleFunc("/path", pathHandler)
	router := Router{}

	// var s Server
	// http.HandleFunc("/home", s.HomeHandler)

	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", router)
	// http.ListenAndServe(":3000", nil)
}
