package database

import (
	"database/sql"
	"fmt"

	"github.com/best2000/rest-api-go/internal/config"
	"github.com/best2000/rest-api-go/internal/logger"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	Db *sql.DB
}

func New(cfg config.Config) *PostgresDb {
	log := logger.Get()

	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.DBName,
	)

	log.Info("db connection string: "+psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresDb{Db: db}
}
