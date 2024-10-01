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

func (m *MongoDBDescriptionRepository) Insert(after *entity.Description) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": after.ProductID}, bson.M{"$push": bson.M{"descriptions": after.Text}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBDescriptionRepository) Update(before, after *entity.Description) error {
	if before.ProductID != after.ProductID {
		return errors.New("product id must be the same (description update)")
	}

	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx,
		bson.M{
			"id":           after.ProductID,
			"descriptions": bson.M{"$elemMatch": bson.M{"$eq": before.Text}},
		},
		bson.M{
			"$set": bson.M{"descriptions.$": after},
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

func (m *MongoDBDescriptionRepository) Delete(before *entity.Description) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx, bson.M{"id": before.ProductID}, bson.M{"$pull": bson.M{"descriptions": before.Text}},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
