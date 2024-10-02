package infra

import (
	"context"
	"errors"
	"time"

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

func (m *MongoDBProductRepository) dtoToEntity(dto *entity.ProductDTO) *entity.Product {
	if dto == nil {
		return nil
	}

	var startDate time.Time
	var endDate *time.Time

	startDate = time.Unix(dto.StartDate/1000000, (dto.StartDate%1000000)*1000000)
	if dto.EndDate <= 0 {
		endDate = nil
	} else {
		e := time.Unix(dto.EndDate/1000000, (dto.EndDate%1000000)*1000000)
		endDate = &e
	}

	return &entity.Product{
		ID:               dto.ID,
		Status:           dto.Status,
		ProductName:      dto.ProductName,
		ProductType:      dto.ProductType,
		SubscriptionType: dto.SubscriptionType,
		StartDate:        startDate,
		EndDate:          endDate,
		UserID:           dto.UserID,
		ParentProductID:  dto.ParentProductID,
		Descriptions:     make([]entity.Description, 0),
		Identifiers:      make([]string, 0),
		Tags:             make([]string, 0),
		Prices:           make([]entity.Price, 0),
		SubProducts:      make([]entity.Product, 0),
	}
}

func (m *MongoDBProductRepository) Insert(afterDTO *entity.ProductDTO) error {
	after := m.dtoToEntity(afterDTO)
	coll := m.db.Collection(UserProductsCollection)

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

func (m *MongoDBProductRepository) Update(beforeDTO, afterDTO *entity.ProductDTO) error {
	before := m.dtoToEntity(beforeDTO)
	after := m.dtoToEntity(afterDTO)

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

func (m *MongoDBProductRepository) Delete(beforeDTO *entity.ProductDTO) error {
	before := m.dtoToEntity(beforeDTO)
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
