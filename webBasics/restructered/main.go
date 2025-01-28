package main

import (
	"fmt"
	"net/http"
)

// Web server port
const portNUmber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("The web server is running on port ", portNUmber)
	http.ListenAndServe(portNUmber, nil)
}
