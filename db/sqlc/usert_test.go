package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	password := util.RandomPassword()
	hashedPassword, err := util.HashPassword(password)

	require.NoError(t, err)

	arg := CreateUserParams{
		Login:           util.RandomUserLogin(),
		FullName:        util.RandomUserName(),
		HashedPasswords: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	assert.NotZero(t, user.CreatedAt)

	assert.Equal(t, user.Login, arg.Login)
	assert.Equal(t, user.FullName, arg.FullName)
	assert.Equal(t, user.HashedPasswords, arg.HashedPasswords)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
