package main

import (
	"log"
	"net/http"

	handler "github.com/brunograsselli/authenticator/http"
	"github.com/brunograsselli/authenticator/postgres"
)

func main() {
	client := postgres.NewClient()

	client.Open()

	defer client.Close()

	handler := handler.NewHandler(client)
	log.Fatal(http.ListenAndServe(":8080", handler.Router()))
}
