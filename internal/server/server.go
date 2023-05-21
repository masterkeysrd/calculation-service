package server

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/common"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/controllers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/middleware"
	"go.uber.org/dig"
)

type Server struct {
	authController        *controllers.AuthController
	userController        *controllers.UserController
	jWTAuthMiddleware     middleware.JWTAuthMiddleware
	operationController   *controllers.OperationController
	calculationController *controllers.CalculationController
	recordController      *controllers.RecordController
}

type ServerParams struct {
	dig.In
	AuthController        *controllers.AuthController
	UserController        *controllers.UserController
	JWTAuthMiddleware     middleware.JWTAuthMiddleware
	OperationController   *controllers.OperationController
	CalculationController *controllers.CalculationController
	RecordController      *controllers.RecordController
}

func NewServer(options ServerParams) *Server {
	return &Server{
		authController:        options.AuthController,
		userController:        options.UserController,
		operationController:   options.OperationController,
		calculationController: options.CalculationController,
		recordController:      options.RecordController,

		jWTAuthMiddleware: options.JWTAuthMiddleware,
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

	authenticated := v1.Group("", s.jWTAuthMiddleware())
	authenticatedControllers := make(map[string]common.Controller)
	authenticatedControllers["users"] = s.userController
	authenticatedControllers["operations"] = s.operationController
	authenticatedControllers["calculations"] = s.calculationController
	authenticatedControllers["records"] = s.recordController

	for path, controller := range authenticatedControllers {
		controller.RegisterRoutes(authenticated.Group(path))
	}
}
