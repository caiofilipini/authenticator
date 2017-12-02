package http

import (
  "net/http"
  "github.com/brunograsselli/authenticator"
  "github.com/brunograsselli/authenticator/crypto"
  "fmt"
)

type Handler struct {
  Client authenticator.Client
}

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
  username, password, ok := r.BasicAuth()

  if ok == false {
    http.Error(w, "Not authorized", 401)
    return
  }

  s := h.Client.CredentialService()

  credential, err := s.Credential(authenticator.Username(username))

  if err != nil {
    http.Error(w, err.Error(), 401)
    return
  }

  auth := &crypto.AuthService{}
  token, err := auth.Authenticate(credential, password)

  if err != nil {
    http.Error(w, "Not authorized", 401)
    return
  }

  fmt.Fprintf(w, "{\"token\":\"%s\"}", token)
}
