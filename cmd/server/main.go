package main

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/balance"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/calculation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/record"
	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/config"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/database"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/random"
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

type RegisterProviders func(*dig.Container) error

func buildContainer() *dig.Container {
	container := dig.New()

	registers := []RegisterProviders{
		jwt.RegisterProviders,
		auth.RegisterProviders,
		user.RegisterProviders,
		config.RegisterProviders,
		record.RegisterProviders,
		server.RegisterProviders,
		balance.RegisterProviders,
		operation.RegisterProviders,
		validator.RegisterProviders,
		middleware.RegisterProviders,
		calculation.RegisterProviders,
		controllers.RegisterProviders,
		database.RegisterProviders,
		persistence.RegisterProviders,
		random.RegisterProviders,
	}

	for _, register := range registers {
		err := register(container)

		if err != nil {
			panic(err)
		}
	}

	return container
}
