package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	http.HandleFunc("/hello", handleHello)

	fmt.Println("Server running at localhost:8000")
	http.ListenAndServe(":8000", nil)
}