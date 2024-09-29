package repository

import (
	"context"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBPriceRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBPriceRepository(conn *database.MongoDBConn) PriceRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBPriceRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBPriceRepository) Update(p *entity.Price) error {
	coll := m.db.Collection(UserProductsCollection)

	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": p.ProductID}, p)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
