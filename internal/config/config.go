package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP httpConfig
		DB   dbConfig
	}
	httpConfig struct {
		Port string
	}
	dbConfig struct {
		Host     string
		Port     string
		Database string
		Username string
		Password string
	}
)

// Init инициализация структуры конфига.
// configFile - Путь до файла конфига yml.
func Init(configFile string) (*Config, error) {
	// read from dotenv
	parseConfig(configFile)

	var cfg Config

	setFromEnv(&cfg)

	if err := unmarshal(&cfg); err != nil {
		log.Fatalf("Cannot unmarshal config: %s", err.Error())
	}

	return &cfg, nil
}

// set parametres from .env.
func setFromEnv(cfg *Config) {
	// database
}

// set parametres from config.yml file.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		log.Fatalf("Error read config: %s", err)
	}

	return nil
}

// parseConfig - парсинг yml файла.
// configPath - название конфига.
func parseConfig(configFile string) error {
	viper.SetConfigFile(fmt.Sprintf("configs/%s", configFile))

	return viper.ReadInConfig()
}
