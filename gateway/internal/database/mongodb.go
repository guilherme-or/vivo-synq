package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConn struct {
	uri    string
	ctx    context.Context
	client *mongo.Client
}

func NewMongoDBConn(uri string) NoSQLConn {
	return &MongoDBConn{uri: uri}
}

func (c *MongoDBConn) Open() error {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(c.uri).SetServerAPIOptions(serverAPI)

	c.ctx = context.TODO()

	// Create a new client and connect to the server
	client, err := mongo.Connect(c.ctx, opts)
	if err != nil {
		return err
	}

	c.client = client

	// Send a ping to confirm a successful connection
	return client.Database("admin").RunCommand(c.ctx, bson.D{{Key: "ping", Value: 1}}).Err()
}

func (c *MongoDBConn) GetClient() interface{} {
	return c.client
}

func (c *MongoDBConn) GetDatabase(name string) any {
	return c.client.Database(name)
}

func (c *MongoDBConn) GetContext() *context.Context {
	return &c.ctx
}

func (c *MongoDBConn) Close() error {
	return c.client.Disconnect(c.ctx)
}
