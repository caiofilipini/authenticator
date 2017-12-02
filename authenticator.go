package authenticator

import (
  "time"
)

type Username string

type Credential struct {
  ID           int       `json:"-"`
  Username     Username  `json:"userName"`
  PasswordHash string    `json:"-"`
  CreatedAt    time.Time `json:createdAt`
  UpdatedAt    time.Time `json:updatedAt`
}

type Client interface {
  CredentialService() CredentialService
}

type CredentialService interface {
  Credential(username Username) (*Credential, error)
}

type Token string

type AuthService interface {
  Authenticate(username string, password string) (*Token, error)
}
