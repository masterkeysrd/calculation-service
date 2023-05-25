package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/common"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/controllers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/middleware"
	"go.uber.org/dig"
)

type Server struct {
	config                *Config
	authController        *controllers.AuthController
	userController        *controllers.UserController
	jWTAuthMiddleware     middleware.JWTAuthMiddleware
	operationController   *controllers.OperationController
	calculationController *controllers.CalculationController
	recordController      *controllers.RecordController
}

type ServerParams struct {
	dig.In
	Config                *Config
	AuthController        *controllers.AuthController
	UserController        *controllers.UserController
	JWTAuthMiddleware     middleware.JWTAuthMiddleware
	OperationController   *controllers.OperationController
	CalculationController *controllers.CalculationController
	RecordController      *controllers.RecordController
}

func NewServer(options ServerParams) *Server {
	return &Server{
		config: options.Config,

		authController:        options.AuthController,
		userController:        options.UserController,
		operationController:   options.OperationController,
		calculationController: options.CalculationController,
		recordController:      options.RecordController,

		jWTAuthMiddleware: options.JWTAuthMiddleware,
	}
}

func (s *Server) Start() error {
	r := s.setup()
	s.registerRoutes(r)

	server := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: r,
	}

	return server.ListenAndServe()
}

func (s *Server) setup() *gin.Engine {
	if s.config.Mode != "" {
		gin.SetMode(s.config.Mode)
	}

	r := gin.Default()

	if s.config.UseCORS {
		s.setupCORS(r)
	}

	return r
}

func (s *Server) setupCORS(r *gin.Engine) {
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		},
	))
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
