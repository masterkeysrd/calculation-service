package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res"
	"github.com/masterkeysrd/calculation-service/internal/server"
	"go.uber.org/dig"
)

func main() {
	container := buildContainer()

	_ = container.Invoke(func(v *validator.Validator) {
		v.RegisterDefaultTranslations()
	})

	err := container.Invoke(func(server *server.Server) {
		server.Start()
	})

	if err != nil {
		panic(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(gin.Default())

	container.Provide(validator.NewValidator)

	container.Provide(server.NewServer)

	container.Provide(res.NewAuthController)
	container.Provide(res.NewUserController)

	container.Provide(user.NewFakeUserRepository)

	container.Provide(auth.NewAuthService)

	container.Provide(user.NewUserFactory)
	container.Provide(user.NewUserService)

	return container
}
