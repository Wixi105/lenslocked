package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wixi105/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, temp string) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	tplPath := filepath.Join("templates", temp)

	tpl, err := views.Parse(tplPath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)

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
	temp := "faq.gohtml"
	executeTemplate(w, temp)
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
