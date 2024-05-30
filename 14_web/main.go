package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> What's up? </h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.HandleFunc("/about", about)

	http.ListenAndServe(":3000", nil)
}
