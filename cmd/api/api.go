package api

import (
	"net/http"
)

type Server struct {
	Addr	string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server"))
}