package main

import (
	"errors"
	"fmt"
	"net/http"
)

// port number for webserver
const portNumber = ":8080"

// create the home page
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page ")
	fmt.Fprintf(w, "%s", fmt.Sprintf("The sum of 2+2 is %d", addTwo(2, 2)))
}

// function to add two numebers
func addTwo(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 0.0
	f, error := divide(x, y)
	alert := "Cannot divide by zero"
	if error != nil {
		fmt.Fprintf(w, alert)
		return
	}
	fmt.Fprintf(w, "%s", fmt.Sprintf("The division of %f by %f is %f", x, y, f))
}

// function to divide two floats numbers and handles the division by zero
func divide(x, y float64) (float64, error) {
	if y != 0 {
		return x / y, nil
	}
	err := errors.New("Cannot Divide by zero")
	return 0, err
}

func main() {
	//create the route to home page, divide and about page
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", Divide)
	//initialize the webserver
	fmt.Println("The web server is running on port ", portNumber)
	http.ListenAndServe(portNumber, nil)
}
