package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConn struct {
	URL      string
	Password string
	User     string
	client   *mongo.Client
	ctx      context.Context
}

func NewMongoConn(url string, password string, user string) NoSQLConn {
	return &MongoConn{
		URL:      url,
		Password: password,
		User:     user,
	}
}

func (m *MongoConn) Open() error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s", m.Password, m.User, m.URL)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, connErr := mongo.Connect(ctx, opts)

	if connErr != nil {
		panic(connErr)
	}

	m.client = client
	m.ctx = ctx

	err := client.Ping(ctx, readpref.Primary())

	return err
}

func (m *MongoConn) GetClient() interface{} {
	return m.client
}

func (m *MongoConn) Close() error {
	return m.client.Disconnect(m.ctx)
}
