package jwt

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewJwtService); err != nil {
		return err
	}

	return nil
}
