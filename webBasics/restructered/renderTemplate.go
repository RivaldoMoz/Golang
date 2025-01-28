package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// This function is used to parse and render html template saved as *.tmpl
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the templete ", err)
		return
	}
}
