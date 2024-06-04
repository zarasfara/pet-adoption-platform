package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/pet-adoption-platform/internal/config"
	"time"
)

const (
	usersTable    = "users"
	petsTable     = "pets"
	sheltersTable = "shelters"
	breedsTable   = "breeds"

	retryCount    = 5
	retryInterval = time.Second * 2
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for i := 0; i < retryCount; i++ {
		db, err = sqlx.Connect("postgres", cfg.DB.String())
		if err == nil {
			return db, nil
		}

		logrus.WithError(err).Warnf("failed to connect to database. Attempt %d of %d. Retrying in %s...", i+1, retryCount, retryInterval)

		time.Sleep(retryInterval)
	}

	return nil, err
}
