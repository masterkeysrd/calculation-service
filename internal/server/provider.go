package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(gin.Default); err != nil {
		return err
	}

	if err := container.Provide(NewServer); err != nil {
		return err
	}

	return nil
}
