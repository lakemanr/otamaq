package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lakemanr/otamaq/db/sqlc"
)

type createRestaurantRequest struct {
	OwnerLogin string `json:"owner_login" binding:"required,min=2,validLogin"`
	Name       string `json:"name" binding:"required,min=2,validName"`
}

func (s *Server) createRestaurant(ctx *gin.Context) {
	var req createRestaurantRequest
	// ShouldBindJSON binds the request body into the struct.
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// If the request body is not a valid JSON or the JSON is not valid for the struct, return an error.
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateRestaurantParams{
		OwnerLogin: req.OwnerLogin,
		Name:       req.Name,
	}
	restaurant, err := s.store.CreateRestaurant(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, restaurant)
}
