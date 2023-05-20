package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res"
	"github.com/masterkeysrd/calculation-service/internal/server"
)

func main() {
	server := buildServer()
	server.RegisterRoutes()
	server.Start()
}

func buildServer() *server.Server {
	validator := validator.NewValidator()
	validator.RegisterDefaultTranslations()

	userFactory := user.NewUserFactory(validator)
	userRepository := user.NewFakeUserRepository()
	userService := user.NewUserService(user.UserServiceOptions{
		CreateUserFactory: userFactory,
		Repository:        userRepository,
	})

	authService := auth.NewAuthService(auth.NewAuthServiceOptions{
		UserService: userService,
	})

	authController := res.NewAuthController(res.NewAuthControllerOptions{
		Service: authService,
	})

	userController := res.NewUserController(res.UserControllerOptions{
		Service: userService,
	})

	return server.NewServer(server.ServerOptions{
		Gin:            gin.Default(),
		AuthController: authController,
		UserController: userController,
	})
}
