package controllers

import "go.uber.org/dig"

func Provide(container *dig.Container) error {
	if err := container.Provide(NewOperationController); err != nil {
		return err
	}

	return nil
}
