package database

import (
	"database/sql"
	"fmt"

	"rest-api/internal/config"
	"rest-api/internal/logger"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connect(cfg config.Config) (*sql.DB, error) {
	log := logger.Logger

	log.Info("connecting to database...")

	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.DBName,
	)

	log.Info("db connection string: " + psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err.Error())
	}

	err = db.Ping()
	
	return db, err
}