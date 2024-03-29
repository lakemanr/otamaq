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
		ID:              util.RandomID(),
		Login:           util.RandomUserLogin(),
		FullName:        util.RandomUserName(),
		HashedPasswords: password,
		CreatedAt:       time.Now().Truncate(time.Second).Local(),
	}
}

func TestCreateUserApi(t *testing.T) {

}
