package database

import (
	"database/sql"
)

type SQLConn interface {
	Open() error
	GetDatabase() *sql.DB
	Close() error
}
