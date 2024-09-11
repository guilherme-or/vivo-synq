package database

import (
	"context"

	"github.com/go-redis/redis"
)

type RedisConn struct {
	opt    *redis.Options
	ctx    context.Context
	client *redis.Client
}

func NewRedisConn(opt *redis.Options) NoSQLConn {
	return &RedisConn{
		opt: opt,
		ctx: context.Background(),
	}
}

func (c *RedisConn) Open() error {
	client := redis.NewClient(c.opt)

	if err := client.Ping().Err(); err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *RedisConn) GetClient() interface{} {
	return c.client
}

func (c *RedisConn) Close() error {
	return c.client.Close()
}
