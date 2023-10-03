package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP httpConfig
	}
	httpConfig struct {
		Port string
	}
)

// Init инициализация структуры конфига.
// configFile - Путь до файла конфига yml.
func Init(configFile string) (*Config, error) {
	return &Config{}, nil
}

// parseConfig - парсинг yml файла.
// configPath - название конфига.
func parseConfig(configFile string) error {
	viper.SetConfigFile(fmt.Sprintf("configs/%s", configFile))

	return viper.ReadInConfig()
}
