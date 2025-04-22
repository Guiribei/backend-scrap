package main

import (
	"fmt"
	"net/http"
	"github.com/Guiribei/backend-scrap/handlers"
)



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cert", handlers.CertHandler)

	fmt.Println("API disponível em http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}