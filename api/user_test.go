package api

import (
	"testing"
	"time"

	db "github.com/lakemanr/otamaq/db/sqlc"
	"github.com/lakemanr/otamaq/util"
)

func createRandomUser() db.User {
	password := util.RandomPassword()

	return db.User{
		Login:           util.RandomUserLogin(),
		FullName:        util.RandomUserName(),
		HashedPasswords: util.HashPassword(password),
		CreatedAt:       time.Now().Truncate(time.Second),
	}
}

func TestCreateUserApi(t *testing.T) {

}
