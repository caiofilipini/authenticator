package main

import (
	"database/sql"
	"log"
	"net/http"

	handler "github.com/brunograsselli/authenticator/http"
	"github.com/brunograsselli/authenticator/postgres"
)

func main() {
	db := mustOpenDB()
	defer db.Close()

	client := postgres.NewClient(db)
	handler := handler.NewHandler(client)

	log.Fatal(http.ListenAndServe(":8080", handler.Router()))
}

func mustOpenDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres dbname=authenticator_development sslmode=disable")
	if err != nil {
		log.Fatal("failed to open postgres DB")
	}
	return db
}
