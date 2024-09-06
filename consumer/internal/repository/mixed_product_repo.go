package repository

import (
	"context"
	"database/sql"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MixedProductRepository struct {
	sqlDB   *sql.DB
	noSqlDB *mongo.Database
	ctx     context.Context
}

// MAKE CONNECTIONS ...
func NewMixedProductRepository(sqlConn *database.PostgreSQLConn, noSqlConn *database.MongoDBConn) ProductRepository {
	sqlDB := sqlConn.GetDatabase()
	noSqlClient := noSqlConn.GetClient().(*mongo.Client)
	noSqlDB := noSqlClient.Database(DatabaseName)

	return &MixedProductRepository{
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
		ctx:     context.TODO(),
	}
}

// MongoDBRepository feature
func (r *MixedProductRepository) Insert(p *entity.Product) error {
	coll := r.noSqlDB.Collection(UserProductsCollection)

	res, err := coll.InsertOne(
		r.ctx,
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

// MongoDBRepository feature
func (r *MixedProductRepository) Update(id int, p *entity.Product) error {
	coll := r.noSqlDB.Collection(UserProductsCollection)

	res, err := coll.UpdateOne(r.ctx, bson.M{"id": id}, p)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

// MongoDBRepository feature
func (r *MixedProductRepository) Delete(id int, productType string) error {
	coll := r.noSqlDB.Collection(UserProductsCollection)

	res, err := coll.DeleteOne(r.ctx, bson.M{"id": id, "product_type": productType})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return ErrNoResult
	}

	return nil
}
