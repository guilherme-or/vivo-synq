package repository

import (
	"context"
	"errors"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DatabaseName = "vivo-synq"
const UserProductsCollection = "user_products"

var ErrNoResult = errors.New("nenhuma linha foi alterada nessa transação")

type MongoDBProductRepository struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

func NewMongoDBProductRepository(conn *database.MongoDBConn) ProductRepository {
	client := conn.GetClient().(*mongo.Client)
	database := conn.GetDatabase("vivo-synq").(*mongo.Database)

	return &MongoDBProductRepository{
		client: client,
		db:     database,
		ctx:    conn.GetContext(),
	}
}

func (m *MongoDBProductRepository) Insert(p *entity.Product) error {
	coll := m.db.Collection(UserProductsCollection)

	res, err := coll.InsertOne(
		*m.ctx,
		p,
	)

	if err != nil {
		return err
	}

	if res.InsertedID == nil {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBProductRepository) Update(id int, p *entity.Product) error {
	coll := m.db.Collection(UserProductsCollection)

	res, err := coll.UpdateOne(*m.ctx, bson.M{"id": id}, p)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

func (m *MongoDBProductRepository) Delete(id int, productType string) error {
	coll := m.db.Collection(UserProductsCollection)

	res, err := coll.DeleteOne(*m.ctx, bson.M{"id": id, "product_type": productType})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNoResult
	}

	return nil
}
