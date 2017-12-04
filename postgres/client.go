package postgres

import (
	"database/sql"

	"github.com/brunograsselli/authenticator"
	_ "github.com/lib/pq"
)

type Client struct {
	db *sql.DB
}

// Ensure Client implements authenticator.CredentialService
// and authenticator.Client
var (
	_ authenticator.Client            = (*Client)(nil)
	_ authenticator.CredentialService = (*Client)(nil)
)

func NewClient(db *sql.DB) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) CredentialService() authenticator.CredentialService { return c }

func (c *Client) Credential(username authenticator.Username) (*authenticator.Credential, error) {
	var credential authenticator.Credential

	row := c.db.QueryRow("SELECT id, username, password_hash, created_at, updated_at FROM credentials WHERE username = $1", username)

	switch err := row.Scan(&credential.ID, &credential.Username, &credential.PasswordHash, &credential.CreatedAt, &credential.UpdatedAt); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &credential, nil
	default:
		return nil, err
	}
}
