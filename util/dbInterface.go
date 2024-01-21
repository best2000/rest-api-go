package util

import (
	"context"
	"database/sql"
)

//using in repo to interchangeable between 'sql.DB' and 'sql.Tx'
type DbExecutor interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type DbQuerist interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}