package main

import (
	"fmt"
	"net/http"
)

// This function handles a request coming from a brownser to :8080/go and returns a string without formation
// to the brownser

func main() {
	http.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		// if the is a error it will printed in the terminal
		if err != nil {
			fmt.Println(err)
		}
		// To display the no of bytes and to display if the is any error
		fmt.Println(fmt.Sprintf("The error in the program is %s", err))
		fmt.Println(fmt.Sprintf("The no of bytes is %d", n))
	})
	http.ListenAndServe(":8080", nil)
}
