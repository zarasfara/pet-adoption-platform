package config

import (
	"fmt"
	"github.com/zarasfara/pet-adoption-platform/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection *gorm.DB

type DatabaseConnection interface {
	Connect() (DBConnection, error)
}

type DBConfig struct {
	Database  string
	Host      string
	Port      string
	Username  string
	Password  string
	SSLMode   string
	Type      string         // Type of the database ("mysql", "postgres", "mssql", etc.).
	Dialector gorm.Dialector // GORM dialector for database configuration.
}

type PostgresConnection struct {
	Config *DBConfig
}

func (config *DBConfig) Connect() (DBConnection, error) {
	db, err := gorm.Open(config.Dialector, &gorm.Config{})

	return db, err
}

// NewConnection creates a new DatabaseConnection based on the given config.
func (config *DBConfig) NewConnection() (DBConnection, error) {
	var dbConnection DatabaseConnection

	switch config.Type {
	case "postgres":
		dbConnection = &PostgresConnection{Config: config}
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}

	// create new connection
	con, err := dbConnection.Connect()
	if err != nil {
		return nil, err
	}

	// AutoMigrate dtos
	err = con.Statement.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return con, nil
}

func (p *PostgresConnection) Connect() (DBConnection, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
	p.Config.Dialector = postgres.Open(fmt.Sprintf(dsn, p.Config.Host, p.Config.Username, p.Config.Password, p.Config.Database, p.Config.Port, p.Config.SSLMode))
	db, err := p.Config.Connect()
	return db, err
}
