package auth

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewAuthService); err != nil {
		return err
	}

	return nil
}
