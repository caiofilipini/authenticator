package main

import (
  "log"
  "github.com/brunograsselli/authenticator/postgres"
  "github.com/brunograsselli/authenticator/crypto"
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

  a := &crypto.AuthService{}

  log.Printf("Testing: %s, %s", c.Username, c.PasswordHash)

  token1, err1 := a.Authenticate(c, "1234567")

  log.Printf("1234567: %s %s", token1, err1)

  token2, err2 := a.Authenticate(c, "123456")

  log.Printf("123456: %s %s", token2, err2)
}
