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

type MongoDBTagRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBTagRepository(conn *database.MongoDBConn) repository.TagRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBTagRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBTagRepository) Insert(after *entity.Tag) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": after.ProductId}, bson.M{"$push": bson.M{"tags": after.Tag}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBTagRepository) Update(before, after *entity.Tag) error {
	if before.ProductId != after.ProductId {
		return errors.New("product id must be the same (tag update)")
	}

	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx,
		bson.M{
			"id":   after.ProductId,
			"tags": bson.M{"$elemMatch": bson.M{"$eq": before.Tag}},
		},
		bson.M{
			"$set": bson.M{"tags.$": after},
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

func (m *MongoDBTagRepository) Delete(before *entity.Tag) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx, bson.M{"id": before.ProductId}, bson.M{"$pull": bson.M{"tags": before.Tag}},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
