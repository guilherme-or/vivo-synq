package infra

import (
	"context"
	"errors"
	"fmt"

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
	fmt.Println("Inserting price: ", after.ID, after.Amount, after.Description, after.ProductID, after.RecurringPeriod, after.Type)
	coll := m.db.Collection(UserProductsCollection)

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": after.ProductID}).Decode(&p); err != nil {
		return err
	}

	p.Prices = append(p.Prices, *after)

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": after.ProductID}, p)
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

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": before.ProductID}).Decode(&p); err != nil {
		return err
	}

	for i, pr := range p.Prices {
		if pr.ID == before.ID {
			p.Prices[i] = *after
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

func (m *MongoDBPriceRepository) Delete(before *entity.Price) error {
	coll := m.db.Collection(UserProductsCollection)

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": before.ProductID}).Decode(&p); err != nil {
		return err
	}

	newPrices := make([]entity.Price, 0)
	for _, pr := range p.Prices {
		if pr.ID != before.ID {
			newPrices = append(newPrices, pr)
		}
	}
	p.Prices = newPrices

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": before.ProductID}, p)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}
