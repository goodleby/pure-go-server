package server

import (
	"log"
	"net/http"

	"example.com/api/env"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Routes map[string]Handler

type Server struct {
	Routes Routes
	Env    env.Variables
}

func Create() (*Server, error) {
	s := Server{
		Routes: Routes{},
	}

	env, err := env.Load()
	if err != nil {
		return nil, err
	}
	s.Env = env

	return &s, nil
}

func (s *Server) AddRoute(path string, handler Handler) {
	http.HandleFunc(path, handler)
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(s.Env.ListenAddress, nil))
}
