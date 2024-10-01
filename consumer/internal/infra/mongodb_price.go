package infra

import (
	"context"
	"errors"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBPriceRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBPriceRepository(conn *database.MongoDBConn) repository.PriceRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBPriceRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBPriceRepository) Insert(after *entity.Price) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": after.ProductID}, bson.M{"$push": bson.M{"prices": after}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBPriceRepository) Update(before, after *entity.Price) error {
	if before.ProductID != after.ProductID {
		return errors.New("product id must be the same (price update)")
	}

	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx,
		bson.M{
			"id":     after.ProductID,
			"prices": bson.M{"$elemMatch": bson.M{"$eq": before.ID}},
		},
		bson.M{
			"$set": bson.M{"prices.$": after},
		},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBPriceRepository) Delete(before *entity.Price) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx, bson.M{"id": before.ProductID}, bson.M{"$pull": bson.M{"prices": bson.M{"id": before.ID}}},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
