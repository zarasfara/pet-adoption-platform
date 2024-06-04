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
		DB   dbConfig
		JWT  jwtConfig
	}
	httpConfig struct {
		Port   string `yaml:"port"`
		AppUrl string
	}
	dbConfig struct {
		Database string
		Host     string
		Port     string
		Username string
		Password string
		SSLMode  string
	}
	jwtConfig struct {
		AccessTokenTTL  time.Duration `yaml:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `yaml:"refreshTokenTTL"`
		SigningToken    string
	}
)

func (cfg dbConfig) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode)
}

// Init initialize Config
func Init(configFile string) (*Config, error) {
	err := parseConfig(configFile)
	if err != nil {
		logrus.Fatalf("cannot unmarshal yml config file: %s", err.Error())
	}

	cfg := new(Config)

	setFromEnv(cfg)
	if err := unmarshal(cfg); err != nil {
		logrus.Fatalf("cannot unmarshal config: %s", err.Error())
	}

	return cfg, nil
}

// Установить параметры из файла dotenv.
func setFromEnv(cfg *Config) {
	cfg.DB = dbConfig{
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}

	cfg.HTTP.AppUrl = os.Getenv("APP_URL")

	cfg.JWT.SigningToken = os.Getenv("SIGNING_TOKEN")
}

// Установить параметры из файла yml.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		logrus.Fatalf("error reading HTTP config: %s", err.Error())
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		logrus.Fatalf("error reading DB config: %s", err.Error())
	}
	if err := viper.UnmarshalKey("auth", &cfg.JWT); err != nil {
		logrus.Fatalf("error reading JWT config: %s", err.Error())
	}

	return nil
}

// parseConfig - парсинг yml конфига.
func parseConfig(configFile string) error {
	viper.SetConfigFile(fmt.Sprintf("configs/%s.yml", configFile))

	return viper.ReadInConfig()
}
