package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/lakemanr/otamaq/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Cannot Load congig:", err)
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
