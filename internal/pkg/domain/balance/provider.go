package balance

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewService); err != nil {
		return err
	}

	return nil
}
