package http

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/brunograsselli/authenticator"
  "log"
)

type Server struct {
  Client authenticator.Client
  handler *Handler
}

func NewServer(client authenticator.Client) (*Server) {
  h := &Handler{Client: client}
  s := &Server{Client: client}
  s.handler = h

  return s
}

func (s *Server) Start() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/authenticate", s.handler.Authenticate).Methods("POST")
  log.Fatal(http.ListenAndServe(":8080", router))
}
