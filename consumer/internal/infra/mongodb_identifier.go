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

type MongoDBIdentifierRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBIdentifierRepository(conn *database.MongoDBConn) repository.IdentifierRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBIdentifierRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBIdentifierRepository) Insert(after *entity.Identifier) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": after.ProductID}, bson.M{"$push": bson.M{"identifiers": after.Identifier}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBIdentifierRepository) Update(before, after *entity.Identifier) error {
	if before.ProductID != after.ProductID {
		return errors.New("product id must be the same (identifier update)")
	}

	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx,
		bson.M{
			"id":          after.ProductID,
			"identifiers": bson.M{"$elemMatch": bson.M{"$eq": before.Identifier}},
		},
		bson.M{
			"$set": bson.M{"identifiers.$": after.Identifier},
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

func (m *MongoDBIdentifierRepository) Delete(before *entity.Identifier) error {
	coll := m.db.Collection(UserProductsCollection)
	res, err := coll.UpdateOne(
		*m.ctx, bson.M{"id": before.ProductID}, bson.M{"$pull": bson.M{"identifiers": before.Identifier}},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
