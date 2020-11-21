package main

import (
	"net/http"
	"log"
)

type Server struct {
	port string
	router *Router
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func NewServer(port string) *Server{
	return &Server{
		port: port,
		router: NewRouter(),
	}

}

func (s *Server) Listen() error {
	log.Print("Server runing in localhost", s.port)
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}