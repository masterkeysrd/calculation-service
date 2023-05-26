package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/loader"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *Config) *gorm.DB {
	databaseLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	dsn := createDatabaseDSN(config)
	pg := postgres.Open(dsn)
	database, err := gorm.Open(pg, &gorm.Config{
		Logger: databaseLogger,
	})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Balance{})
	database.AutoMigrate(&models.Operation{})
	database.AutoMigrate(&models.Record{})

	err = loader.LoadDefaultOperations(database)
	if err != nil {
		panic(err)
	}

	return database
}

func createDatabaseDSN(config *Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.DBName, config.Password)
}
