package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/best2000/rest-api-go/config"
	_ "github.com/lib/pq"
)

type PostgresDb struct {
	Db *sql.DB
}

func NewPostgresDatabase(cfg config.Config) *PostgresDb {
	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.DBName,
	)

	slog.Info("db connection string: "+psqlInfo)

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
