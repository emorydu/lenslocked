package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// bio := `<script>alert("Haha, you have been h4x0r3d!"); </script>`
	bio := `&lt;script&gt;alert(&quot;Hi!&quot;);&lt;/script&gt;`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1><p>Bio:"+bio+"</p>")
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
