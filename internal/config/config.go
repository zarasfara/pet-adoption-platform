package config

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	Config struct {
		HTTP httpConfig
		DB   DBConfig
		JWT  JWTConfig
	}
	httpConfig struct {
		Port string
	}
	DBConfig struct {
		Database string
		Host     string
		Port     string
		Username string
		Password string
		SSLMode  string
	}
	JWTConfig struct {
		RefreshTokenTTL time.Duration
		SigningToken    string
	}
)

// Init инициализация структуры конфига.
func Init(configFile string) (*Config, error) {
	// Чтение из dotenv
	err := parseConfig(configFile)
	if err != nil {
		logrus.Fatalf("Cannot unmarshal yml config file: %s", err.Error())
	}

	cfg := new(Config) // Создаем экземпляр структуры Config

	setFromEnv(cfg)
	if err := unmarshal(cfg); err != nil {
		logrus.Fatalf("Cannot unmarshal config: %s", err.Error())
	}

	return cfg, nil
}

// Установить параметры из файла dotenv.
func setFromEnv(cfg *Config) {
	// database
	cfg.DB.Database = os.Getenv("DB_DATABASE")
	cfg.DB.Username = os.Getenv("DB_USERNAME")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")

	// authentication
	cfg.JWT.SigningToken = os.Getenv("SIGNING_TOKEN")
}

// Установить параметры из файла yml.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		logrus.Fatalf("Error read config: %s", err.Error())
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		logrus.Fatalf("Error read config: %s", err.Error())
	}
	if err := viper.UnmarshalKey("auth", &cfg.JWT); err != nil {
		logrus.Fatalf("Error read config: %s", err.Error())
	}

	return nil
}

// parseConfig - парсинг yml конфига.
func parseConfig(configFile string) error {
	viper.SetConfigFile(fmt.Sprintf("configs/%s.yml", configFile))

	return viper.ReadInConfig()
}
