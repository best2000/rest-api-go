package repo

import (
	"database/sql"
)

type CustomRepo struct {
	Db *sql.DB
}