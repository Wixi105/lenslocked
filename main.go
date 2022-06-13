package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1> Welcome to my new website</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1> Contact me</h1> <p>To get in touch, email me at <a href=\"mailto:ericsampson109@gmail.com\">ericsampson@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Frequently Asked Questions</h1>
	
	<ul>
		<li>Can we get a Free Trial? Yes, free trials are available</li>
		<li>How many people can signup for this event? 300 people per ticket</li>
		<li>Is this difficult? No, not at all. If you do your best, you will be well rewarded.</li>
	</ul>
	
	`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	fmt.Println("Starting server on port: 3000")
	http.ListenAndServe(":3000", r)
}
