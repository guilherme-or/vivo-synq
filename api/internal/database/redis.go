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

func (c *RedisConn) Open() error {
	opt := &redis.Options{
		Addr:      c.Addr,
		Password:  c.Password,
		DB:        c.DB,
		TLSConfig: nil,
	}

	if c.URL != "" {
		parsedURL, err := redis.ParseURL(c.URL)
		if err != nil {
			return err
		}
		opt = parsedURL
	}

	c.client = redis.NewClient(opt)

	_, err := c.client.Ping().Result()
	return err
}

func (c *RedisConn) GetClient() interface{} {
	return c.client
}

func (c *RedisConn) Close() error {
	return c.client.Close()
}
