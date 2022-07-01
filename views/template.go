package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmltpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	err := t.htmltpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing Template: %v", err)
		http.Error(w, "There was a problem executing the template", http.StatusInternalServerError)
		return
	}
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("Parsing template failed: %v", err)
	}

	return Template{tpl}, nil
}
