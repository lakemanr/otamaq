package main

import (
	"database/sql"
	"log"

	"github.com/lakemanr/otamaq/api"
	db "github.com/lakemanr/otamaq/db/sqlc"
	"github.com/lakemanr/otamaq/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	s := api.NewServer(db.NewStore(conn))
	s.Start(config.ServerAddress)
}
