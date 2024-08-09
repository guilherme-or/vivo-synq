package database

import (
	"github.com/go-redis/redis"
)

type RedisConn struct {
	Addr     string
	Password string
	DB       int
	// Example: redis://<user>:<pass>@localhost:6379/<db>
	URL    string
	client *redis.Client
}

func NewRedisConn(addr, password string, db int) NoSQLConn {
	return &RedisConn{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
}

func NewRedisConnWithURL(url string) NoSQLConn {
	return &RedisConn{
		URL: url,
	}
}

func (r *RedisConn) Open() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:      r.Addr,
		Password:  r.Password,
		DB:        r.DB,
		TLSConfig: nil,
	})

	_, err := r.client.Ping().Result()
	return err
}

func (r *RedisConn) GetClient() interface{} {
	return r.client
}

func (r *RedisConn) Close() error {
	return r.client.Close()
}
