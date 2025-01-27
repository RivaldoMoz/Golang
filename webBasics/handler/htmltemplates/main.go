package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//Working with html templates

// port number for the http server
const portNumber = ":8080"

// This function is used to display the home page
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// This function is used to display the about page
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// This function is used to parse an template file and render it
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template ", err)
	}
	return
}

func main() {
	//This is a handler to process the request to the home page func
	http.HandleFunc("/", home)
	//This is a handler to process the request to the about page func
	http.HandleFunc("/about", about)
	//Run webserver on port 8080
	fmt.Println(fmt.Sprintf("The server is running on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
