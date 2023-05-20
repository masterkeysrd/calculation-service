package calculation

import "go.uber.org/dig"

func Provide(container *dig.Container) error {
	if err := container.Provide(NewService); err != nil {
		return err
	}

	return nil
}
