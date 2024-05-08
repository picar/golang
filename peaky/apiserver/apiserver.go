package apiserver

import (
	"log"
	"net/http"
)

type Server struct {
	Address string
}

func NewServer(s Server) Server {
	log.Printf("%+v\n", s)
	return s
}

func (s *Server) AddRoutes() error {
	return nil
}

func (s *Server) Start() {
	err := http.ListenAndServe(s.Address, nil)
	if err != nil {
		log.Println(err)
	}
}
