package postgres

import (
  "github.com/brunograsselli/authenticator"
  "database/sql"
)

// Ensure CredentialService implements authenticator.CredentialService
var _ authenticator.CredentialService = &CredentialService{}

type CredentialService struct {
  client *Client
}

func (c *CredentialService) Credential(username authenticator.Username) (*authenticator.Credential, error) {
  credential := authenticator.Credential{}

  row := c.client.db.QueryRow("SELECT id, username, password_hash, created_at, updated_at FROM credentials WHERE username = $1", username)

  switch err := row.Scan(&credential.ID, &credential.Username, &credential.PasswordHash, &credential.CreatedAt, &credential.UpdatedAt); err {
  case sql.ErrNoRows:
    return nil, nil
  case nil:
    return &credential, nil
  default:
    return nil, err
  }
}
