package main

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/balance"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/calculation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/config"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/controllers"
	"github.com/masterkeysrd/calculation-service/internal/pkg/web/res/middleware"
	"github.com/masterkeysrd/calculation-service/internal/server"
	"go.uber.org/dig"
)

func main() {
	config.LoadConfig()

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

	jwt.RegisterProviders(container)
	auth.RegisterProviders(container)
	user.RegisterProviders(container)
	config.RegisterProviders(container)
	record.RegisterProviders(container)
	server.RegisterProviders(container)
	balance.RegisterProviders(container)
	operation.RegisterProviders(container)
	validator.RegisterProviders(container)
	middleware.RegisterProviders(container)
	calculation.RegisterProviders(container)
	controllers.RegisterProviders(container)

	return container
}
