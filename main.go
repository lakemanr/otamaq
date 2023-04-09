package main

import (
	"database/sql"
	"log"

	"github.com/lakemanr/otamaq/api"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/otamaq?sslmode=disable"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	s := api.NewServer(conn)
	s.Start("localhost:8080")
}
