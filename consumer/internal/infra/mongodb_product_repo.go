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

const (
	DatabaseName           = "vivo-synq"
	UserProductsCollection = "user_products"
)

var ErrNoResult = errors.New("nenhuma linha foi alterada nessa transação")

type MongoDBProductRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBProductRepository(conn *database.MongoDBConn) repository.ProductRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	coll := database.Collection(UserProductsCollection)
	coll.DeleteMany(context.Background(), bson.M{})

	return &MongoDBProductRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBProductRepository) Insert(after *entity.Product) error {
	coll := m.db.Collection(UserProductsCollection)

	after.Descriptions = make([]entity.Description, 0)
	after.Identifiers = make([]string, 0)
	after.Tags = make([]string, 0)
	after.Prices = make([]entity.Price, 0)
	after.SubProducts = make([]entity.Product, 0)

	// Insert on parent product
	if after.ParentProductID != nil && *after.ParentProductID > 0 {
		var parent entity.Product
		if err := coll.FindOne(*m.ctx, bson.M{"id": after.ParentProductID}).Decode(&parent); err != nil {
			return err
		}
		parent.SubProducts = append(parent.SubProducts, *after)
		coll.ReplaceOne(*m.ctx, bson.M{"id": after.ParentProductID}, parent)
	}

	res, err := coll.InsertOne(
		*m.ctx,
		after,
	)

	if err != nil {
		return err
	}

	if res.InsertedID == nil {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBProductRepository) Update(before, after *entity.Product) error {
	if before.ID != after.ID {
		return errors.New("product id must be the same (tag update)")
	}
	coll := m.db.Collection(UserProductsCollection)

	var p entity.Product
	if err := coll.FindOne(*m.ctx, bson.M{"id": before.ID}).Decode(&p); err != nil {
		return err
	}

	after.Descriptions = p.Descriptions
	after.Identifiers = p.Identifiers
	after.Tags = p.Tags
	after.Prices = p.Prices
	after.SubProducts = p.SubProducts

	// Update on parent product
	if after.ParentProductID != nil && *after.ParentProductID > 0 {
		var parent entity.Product
		if err := coll.FindOne(*m.ctx, bson.M{"id": after.ParentProductID}).Decode(&parent); err != nil {
			return err
		}
		for i, sub := range parent.SubProducts {
			if sub.ID == before.ID {
				parent.SubProducts[i] = *after
				break
			}
		}
		coll.ReplaceOne(*m.ctx, bson.M{"id": after.ParentProductID}, parent)
	}

	res, err := coll.ReplaceOne(*m.ctx, bson.M{"id": before.ID}, after)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBProductRepository) Delete(before *entity.Product) error {
	coll := m.db.Collection(UserProductsCollection)

	// Delete on parent product
	if before.ParentProductID != nil && *before.ParentProductID > 0 {
		var parent entity.Product
		if err := coll.FindOne(*m.ctx, bson.M{"id": before.ParentProductID}).Decode(&parent); err != nil {
			return err
		}
		newSubProducts := make([]entity.Product, 0)
		for _, sub := range parent.SubProducts {
			if sub.ID != before.ID {
				newSubProducts = append(newSubProducts, sub)
			}
		}
		parent.SubProducts = newSubProducts
		coll.ReplaceOne(*m.ctx, bson.M{"id": before.ParentProductID}, parent)
	}

	res, err := coll.DeleteOne(*m.ctx, bson.M{"id": before.ID})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNoResult
	}

	return nil
}
