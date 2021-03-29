package main

import "net/http"


type Server struct {
	port string
	router *Router
}

func newServer(port string) *Server{
	return &Server{
		port: port,
		router: newRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}