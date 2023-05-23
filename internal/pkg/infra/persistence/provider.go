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
	// Register userRepository
	if err := container.Provide(repositories.NewUserRepository); err != nil {
		return err
	}

	return nil
}
