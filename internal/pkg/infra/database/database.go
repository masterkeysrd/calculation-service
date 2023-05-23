package database

import (
	"fmt"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *Config) *gorm.DB {
	dsn := createDatabaseDSN(config)
	pg := postgres.Open(dsn)
	database, err := gorm.Open(pg, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Balance{})
	return database
}

func createDatabaseDSN(config *Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.DBName, config.Password)
}
