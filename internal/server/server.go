package server

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/auth"
)

type Server struct {
	gin            *gin.Engine
	authController *auth.Controller
}

func NewServer() *Server {
	return &Server{
		gin:            gin.Default(),
		authController: auth.NewController(),
	}
}

func (s *Server) RegisterRoutes() {
	api := s.gin.Group("/api")
	v1 := api.Group("/v1")

	s.authController.RegisterRoutes(v1.Group("/auth"))
}

func (s *Server) Start() error {
	return s.gin.Run()
}
