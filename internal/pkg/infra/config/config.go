package config

import (
	"fmt"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/database"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/random"
	"github.com/masterkeysrd/calculation-service/internal/server"
)

type Config struct {
	Server   *server.Config   `mapstructure:"server"`
	JWT      *jwt.Config      `mapstructure:"jwt"`
	Database *database.Config `mapstructure:"database"`
	Services *ConfigServices  `mapstructure:"services"`
}

type ConfigServices struct {
	Random *random.Config `mapstructure:"random"`
}

func LoadConfig() {
	profile := config.Getenv("APP_ENV", "local")
	configPath := config.Getenv("APP_CONFIG_PATH", "../../config")

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	file := fmt.Sprintf("%s/%s.yml", configPath, profile)
	err := config.LoadFiles(file)
	if err != nil {
		panic(err)
	}
}

func Get() *Config {
	cfg := &Config{}
	config.Decode(cfg)
	return cfg
}

func JWT(config *Config) *jwt.Config {
	return config.JWT
}

func Database(config *Config) *database.Config {
	return config.Database
}

func ServicesRandom(config *Config) *random.Config {
	return config.Services.Random
}

func Server(config *Config) *server.Config {
	return config.Server
}
