package infra

import (
	"context"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBDescriptionRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBDescriptionRepository(conn *database.MongoDBConn) repository.DescriptionRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBDescriptionRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBDescriptionRepository) Update(d *entity.Description) error {
	coll := m.db.Collection(UserProductsCollection)

	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": d.ProductID}, d)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
