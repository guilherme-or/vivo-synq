package database

import "context"

type DatabaseConn interface {
	Open() error
	Close() error
	GetClient() any
	GetContext() context.Context
}
