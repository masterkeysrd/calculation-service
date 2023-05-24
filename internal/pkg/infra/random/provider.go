package random

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewClient); err != nil {
		return err
	}

	return nil
}
