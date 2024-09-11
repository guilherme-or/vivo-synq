package database

import (
	"database/sql"
)

type SQLConn interface {
	Open() error
	GetDatabase() *sql.DB
	Close() error
}

type NoSQLConn interface {
	Open() error
	GetClient() interface{}
	Close() error
}
