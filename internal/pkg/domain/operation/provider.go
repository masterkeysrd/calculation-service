package operation

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewRepository); err != nil {
		return err
	}

	if err := container.Provide(NewOperationService); err != nil {
		return err
	}

	return nil
}
