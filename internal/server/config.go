package server

type Config struct {
	Mode    string `mapstructure:"mode"`
	Port    string `mapstructure:"port"`
	UseCORS bool   `mapstructure:"useCORS"`
}
