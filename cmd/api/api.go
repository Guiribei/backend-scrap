package api

import (
	"net/http"
)

type Server struct {
	Addr	string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index page!"))
			return
		case "/dale":
			w.Write([]byte("Daaale page!"))
			return
		}
	default:
		w.Write([]byte("404"))
		return
	}

}