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

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": after.ProductID}).Decode(&p); err != nil {
		return err
	}

	p.Descriptions = append(p.Descriptions, *after)

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": after.ProductID}, p)
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

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": after.ProductID}).Decode(&p); err != nil {
		return err
	}

	for i, d := range p.Descriptions {
		if d.ID == before.ID {
			p.Descriptions[i] = *after
			break
		}
	}

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": after.ProductID}, p)
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

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": before.ProductID}).Decode(&p); err != nil {
		return err
	}

	newDescriptions := make([]entity.Description, 0)
	for _, d := range p.Descriptions {
		if d.ID != before.ID {
			p.Descriptions = append(p.Descriptions, *before)
		}
	}

	p.Descriptions = newDescriptions

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": before.ProductID}, p)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
