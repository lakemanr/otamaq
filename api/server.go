package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validName", validName)
		v.RegisterValidation("validLogin", validLogin)
	}

	s.router.POST("/users", s.createUser)
	s.router.POST("/restaurants", s.createRestaurant)

	return s
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
