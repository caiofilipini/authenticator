package crypto

import (
  "github.com/brunograsselli/authenticator"
  "golang.org/x/crypto/bcrypt"
  "errors"
)

var _ authenticator.AuthService = &AuthService{}

type AuthService struct {
}

func (a *AuthService) Authenticate(credential *authenticator.Credential, password string) (authenticator.Token, error) {
  if a.samePassword(credential.PasswordHash, password) {
    return "fake token", nil
  } else {
    return "", errors.New("invalid credentials")
  }
}

func (a *AuthService) samePassword(hash string, password string) bool {
  byteHash := []byte(hash)
  plainPwd := []byte(password)
  err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

  return err == nil
}
