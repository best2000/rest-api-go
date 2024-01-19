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
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password, cfg.Db.DBName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("ping to db.")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresDb{Db: db}
}
