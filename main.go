package main

import (
	"database/sql"
	"log"

	"github.com/lakemanr/otamaq/api"
	"github.com/lakemanr/otamaq/util"
	_ "github.com/lib/pq"
)

func main() {

	congig, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(congig.DbDriver, congig.DbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	s := api.NewServer(conn)
	s.Start(congig.ServerAddress)
}
