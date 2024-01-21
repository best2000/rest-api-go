package database

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/best2000/rest-api-go/config"
)

type PostgresDb struct {
	Db *sql.DB
}

func NewPostgresDatabase(cfg config.Config) *PostgresDb {
	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.DBName,
	)

	fmt.Println("db connection string: "+psqlInfo)

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
