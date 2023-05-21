package config

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(GetConfig); err != nil {
		return err
	}

	if err := container.Provide(GetJWTConfig); err != nil {
		return err
	}

	return nil
}
