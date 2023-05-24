package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/database"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/random"
)

type Config struct {
	JWT      *jwt.Config      `mapstructure:"jwt"`
	Database *database.Config `mapstructure:"database"`
	Services *ConfigServices  `mapstructure:"services"`
}

type ConfigServices struct {
	Random *random.Config `mapstructure:"random"`
}

func LoadConfig() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	// TODO: Change the path to a ENV variable
	err := config.LoadFiles("../../config/server.yml")
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	cfg := &Config{}
	config.Decode(cfg)
	return cfg
}

func GetJWTConfig(config *Config) *jwt.Config {
	return config.JWT
}

func GetDatabaseConfig(config *Config) *database.Config {
	return config.Database
}

func GetRandomConfig(config *Config) *random.Config {
	return config.Services.Random
}
