package middleware

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(JWTAuthMiddlewareFactory); err != nil {
		return err
	}

	return nil
}
