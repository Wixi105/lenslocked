package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, temp string) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	tplPath := filepath.Join("templates", temp)
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Failed to load template: %v", err)
		http.Error(w, "There was a problem loading the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing Template: %v", err)
		http.Error(w, "There was a problem executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	temp := "home.gohtml"
	executeTemplate(w, temp)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	temp := "contact.gohtml"
	executeTemplate(w, temp)

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

func articleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	articleByID := chi.URLParam(r, "id")

	switch articleByID {
	case "1":
		fmt.Fprintf(w, `
		<h1>This is the first article</h1>
		<p>Article %v</p>
		`, articleByID)
	case "2":
		fmt.Fprintf(w, `
			<h1>This is the second article</h1>
			<p>Article %v</p>
			`, articleByID)
	}

}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/article/{id}", articleHandler)

	fmt.Println("Starting server on port: 3000")
	http.ListenAndServe(":3000", r)
}
