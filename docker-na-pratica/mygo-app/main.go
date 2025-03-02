package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		fmt.Printf("Starting server on port %s\n", args[1])

		http.ListenAndServe(":"+args[1], http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		}))	
	} else {

		fmt.Println("Starting server on port 8080")

		http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		}))
	}
}
