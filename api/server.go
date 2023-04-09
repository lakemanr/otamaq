package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Server {
	s := &Server{
		db:     db,
		router: gin.Default(),
	}

	s.router.POST("/restaurants", s.createRestaurant)

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
