package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/lakemanr/otamaq/db/sqlc"
	"github.com/lakemanr/otamaq/util"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Login    string `json:"login" binding:"required,min=2,validLogin"`
	FullName string `json:"full_name" binding:"required,min=2,validName"`
	Password string `json:"password" binding:"required,min=8"`
}

type createUserResponse struct {
	ID        int32     `json:"id"`
	Login     string    `json:"login"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	arg := db.CreateUserParams{
		Login:           req.Login,
		FullName:        req.FullName,
		HashedPasswords: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("user with login %s already exists", arg.Login)})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	rsp := createUserResponse{
		ID:        user.ID,
		Login:     user.Login,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)

}
