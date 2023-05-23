package database

import (
	"fmt"

	"go.uber.org/dig"
)

func RegisterProviders(container *dig.Container) error {
	fmt.Println("Registering database providers")
	if err := container.Provide(NewDatabase); err != nil {
		return err
	}

	return nil
}
