package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"go.tmp/echo/internal/net/http"
)

func main() {
	// addr := os.Getenv("DATABASE_URL")

	// assume this has been constructed
	var connSQL *sqlx.DB

	log.Fatal(http.NewServer(connSQL).Start(":8080"))
}
