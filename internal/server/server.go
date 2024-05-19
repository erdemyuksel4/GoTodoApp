package server

import (
	"net/http"
)


func (s *Server) StartServer(addr string) {
	mux := http.NewServeMux()


	http.ListenAndServe(addr, mux)
}
