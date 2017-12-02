package crypto

import (
  "github.com/brunograsselli/authenticator"
  "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
  "errors"
)

var _ authenticator.AuthService = &AuthService{}

type AuthService struct {
}

func (a *AuthService) Authenticate(credential *authenticator.Credential, password string) (authenticator.Token, error) {
  if a.samePassword(credential.PasswordHash, password) {
    token, err := a.jwtTokenFor(credential)

    if err != nil {
      return "", err
    }

    return authenticator.Token(token), nil
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

func (a *AuthService) jwtTokenFor(credential *authenticator.Credential) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": credential.Username,
  })

  tokenString, err := token.SignedString([]byte("secret"))

  if err != nil {
    return "", err
  }

  return tokenString, nil
}
