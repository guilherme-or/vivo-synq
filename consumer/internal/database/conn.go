package database

import (
	"context"
	"database/sql"
)

type SQLConn interface {
	Open() error
	GetDatabase() *sql.DB
	Close() error
}

type NoSQLConn interface {
	Open() error
	GetClient() any
	GetDatabase(name string) any
	GetContext() *context.Context
	Close() error
}