package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lakemanr/otamaq/db/sqlc"
)

type createUserRequest struct {
	Login    string `json:"login" binding:"required,min=2,validLogin"`
	FullName string `json:"full_name" binding:"required,min=2,validName"`
	Password string `json:"password" binding:"required,min=8"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	arg := db.CreateUserParams{
		Login:           req.Login,
		FullName:        req.FullName,
		HashedPasswords: req.Password,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)

}
