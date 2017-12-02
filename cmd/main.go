package main

import (
  "log"
  "github.com/brunograsselli/authenticator/postgres"
)

func main() {
  client := postgres.NewClient()

  client.Open()

  defer client.Close()

  s := client.CredentialService()

  c, error := s.Credential("user")

  if error != nil {
    log.Fatal(error)
  }

  log.Printf("Testing: %s, %s", c.Username, c.PasswordHash)
}
