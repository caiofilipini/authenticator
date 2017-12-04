package main

import (
  "github.com/brunograsselli/authenticator/postgres"
  "github.com/brunograsselli/authenticator/http"
)

func main() {
  client := postgres.NewClient()

  client.Open()

  defer client.Close()

  server := http.NewServer(client)

  server.Start()
}
