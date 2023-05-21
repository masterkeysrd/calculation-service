package controllers

import "go.uber.org/dig"

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(NewUserController); err != nil {
		return err
	}

	if err := container.Provide(NewAuthController); err != nil {
		return err
	}

	if err := container.Provide(NewOperationController); err != nil {
		return err
	}

	if err := container.Provide(NewCalculationController); err != nil {
		return err
	}

	if err := container.Provide(NewRecordController); err != nil {
		return err
	}

	return nil
}
