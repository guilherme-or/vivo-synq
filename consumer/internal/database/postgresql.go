package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgreSQLConn struct {
	DSN      string
	database *sql.DB
}

func NewPostgreSQLConn(dsn string) SQLConn {
	return &PostgreSQLConn{
		DSN: dsn,
	}
}

func (c *PostgreSQLConn) Open() error {
	db, err := sql.Open("postgres", c.DSN)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	c.database = db

	return nil
}

func (c *PostgreSQLConn) GetDatabase() *sql.DB {
	return c.database
}

func (c *PostgreSQLConn) Close() error {
	return c.database.Close()
}
