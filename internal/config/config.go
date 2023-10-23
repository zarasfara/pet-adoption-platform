package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP httpConfig
		DB   DBConfig
	}
	httpConfig struct {
		Port string
	}
)

// Init инициализация структуры конфига.
// configFile - Путь до файла конфига yml.
func Init(configFile string) (*Config, error) {
	// Чтение из dotenv
	err := parseConfig(configFile)
	if err != nil {
		log.Fatalf("Cannot unmarshal yml config file: %s", err.Error())
	}

	cfg := &Config{} // Создаем экземпляр структуры Config

	setFromEnv(cfg)
	if err := unmarshal(cfg); err != nil {
		log.Fatalf("Cannot unmarshal config: %s", err.Error())
	}

	return cfg, nil
}

// set parameters from .env.
func setFromEnv(cfg *Config) {
	// database
	cfg.DB.Database = os.Getenv("DB_DATABASE")
	cfg.DB.Username = os.Getenv("DB_USERNAME")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.Type = os.Getenv("DB_TYPE")
}

// set parameters from config.yml file.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		log.Fatalf("Error read config: %s", err)
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		log.Fatalf("Error read config: %s", err)
	}

	return nil
}

// parseConfig - парсинг yml файла.
// configFile - название конфига.
func parseConfig(configFile string) error {
	viper.SetConfigFile(fmt.Sprintf("configs/%s.yml", configFile))

	return viper.ReadInConfig()
}
