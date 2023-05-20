package server

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res"
	"go.uber.org/dig"
)

type Server struct {
	authController *res.AuthController
	userController *res.UserController
}

type ServerParams struct {
	dig.In
	AuthController *res.AuthController
	UserController *res.UserController
}

func NewServer(options ServerParams) *Server {
	return &Server{
		authController: options.AuthController,
		userController: options.UserController,
	}
}

func (s *Server) Start() error {
	r := gin.Default()
	s.registerRoutes(r)

	return r.Run(":8080")
}

func (s *Server) registerRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	s.authController.RegisterRoutes(v1.Group("/auth"))
	s.userController.RegisterRoutes(v1.Group("/users"))
}
