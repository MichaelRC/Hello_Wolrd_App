package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//renderTemplate will render the HTML files stored in /templates/
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
