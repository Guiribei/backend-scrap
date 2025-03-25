package main

import (
	"log"
	"net/http"

	"github.com/Guiribei/backend-scrap/cmd/api"
)

func main() {
	s:= &api.Server{Addr:":8080"}
	if err := http.ListenAndServe(s.Addr, s); err != nil {
		log.Fatal(http.ListenAndServe(s.Addr, s))
	}
}