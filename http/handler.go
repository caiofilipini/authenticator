package http

import (
	"fmt"
	"net/http"

	"github.com/brunograsselli/authenticator"
	"github.com/brunograsselli/authenticator/crypto"
	"github.com/gorilla/mux"
)

type Handler struct {
	authClient authenticator.Client
}

func NewHandler(authClient authenticator.Client) Handler {
	return Handler{
		authClient: authClient,
	}
}

func (h *Handler) Router() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/authenticate", h.authenticate).Methods("POST")
	return router
}

func (h *Handler) authenticate(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if ok == false {
		http.Error(w, "Not authorized", 401)
		return
	}

	s := h.authClient.CredentialService()

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
