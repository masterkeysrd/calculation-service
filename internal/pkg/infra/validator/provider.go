package validator

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewValidator); err != nil {
		return err
	}

	return nil
}
