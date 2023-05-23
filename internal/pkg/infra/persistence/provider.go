package persistence

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/repositories"
	"go.uber.org/dig"
)

func RegisterProviders(container *dig.Container) error {
	// Register repositories
	if err := registerRepositories(container); err != nil {
		return err
	}

	return nil
}

func registerRepositories(container *dig.Container) error {
	providers := []interface{}{
		repositories.NewUserRepository,
		repositories.NewOperationRepository,
		repositories.NewRecordRepository,
		repositories.NewBalanceRepository,
	}

	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			return err
		}
	}

	return nil
}
