package postgres

import (
  "database/sql"

  "github.com/brunograsselli/authenticator"
  _ "github.com/lib/pq"
)

type Client struct {
  db *sql.DB
  credentialService CredentialService
}

func NewClient() *Client {
	c := &Client{}
	c.credentialService.client = c
	return c
}

func (c *Client) Open() error {
  var err error
  c.db, err = sql.Open("postgres", "user=postgres dbname=authenticator_development sslmode=disable")

  if err != nil {
    return err
  }

  return nil
}

func (c *Client) Close() error {
  err := c.db.Close()

  if err != nil {
    return err
  }

  return nil
}

func (c *Client) CredentialService() authenticator.CredentialService { return &c.credentialService }
