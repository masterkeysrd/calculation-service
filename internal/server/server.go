package server

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res"
)

type Server struct {
	gin            *gin.Engine
	authController *res.AuthController
	userController *res.UserController
}

type ServerOptions struct {
	Gin            *gin.Engine
	AuthController *res.AuthController
	UserController *res.UserController
}

func NewServer(options ServerOptions) *Server {
	return &Server{
		gin:            options.Gin,
		authController: options.AuthController,
		userController: options.UserController,
	}
}

func (s *Server) RegisterRoutes() {
	api := s.gin.Group("/api")
	v1 := api.Group("/v1")

	s.authController.RegisterRoutes(v1.Group("/auth"))
	s.userController.RegisterRoutes(v1.Group("/users"))
}

func (s *Server) Start() error {
	return s.gin.Run(":8080")
}
