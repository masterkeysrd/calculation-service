package config

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(Get); err != nil {
		return err
	}

	if err := container.Provide(JWT); err != nil {
		return err
	}

	if err := container.Provide(Database); err != nil {
		return err
	}

	if err := container.Provide(ServicesRandom); err != nil {
		return err
	}

	if err := container.Provide(Server); err != nil {
		return err
	}

	return nil
}
