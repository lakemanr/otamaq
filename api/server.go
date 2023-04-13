package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/lakemanr/otamaq/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	s := &Server{
		store:  store,
		router: gin.Default(),
	}

	s.router.POST("/restaurants", s.createRestaurant)

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
