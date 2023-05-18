package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer() *Server {
	return &Server{
		gin: gin.Default(),
	}
}

func (s *Server) Start() error {
	return s.gin.Run()
}
