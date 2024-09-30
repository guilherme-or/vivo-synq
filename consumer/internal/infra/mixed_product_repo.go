package infra

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	queryGetCompleteProduct = `
		SELECT * FROM get_complete_product($1)
	`
)

type MixedProductRepository struct {
	sqlDB   *sql.DB
	noSqlDB *mongo.Database
	ctx     context.Context
}

func NewMixedProductRepository(sqlConn *database.PostgreSQLConn, noSqlConn *database.MongoDBConn) repository.ProductRepository {
	sqlDB := sqlConn.GetDatabase()
	noSqlClient := noSqlConn.GetClient().(*mongo.Client)
	noSqlDB := noSqlClient.Database(DatabaseName)

	coll := noSqlDB.Collection(UserProductsCollection)
	coll.DeleteMany(context.Background(), bson.M{})

	return &MixedProductRepository{
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
		ctx:     context.TODO(),
	}
}

func (r *MixedProductRepository) tryCompleteProduct(incomplete *entity.Product) *entity.Product {
	rows, err := r.sqlDB.Query(queryGetCompleteProduct, incomplete.ID)
	if err != nil {
		return incomplete
	}
	defer rows.Close()

	// Processar os resultados em JSON ([]byte)
	var jsonResult []byte
	for rows.Next() {
		if err := rows.Scan(&jsonResult); err != nil {
			return incomplete
		}
	}

	// Deserializar o JSON para a estrutura Product
	var product entity.Product
	if err := json.Unmarshal(jsonResult, &product); err != nil {
		return incomplete
	}

	return &product
}

func (r *MixedProductRepository) Insert(p *entity.Product) error {
	complete := r.tryCompleteProduct(p)
	coll := r.noSqlDB.Collection(UserProductsCollection)

	res, err := coll.InsertOne(
		r.ctx,
		complete,
	)

	if err != nil {
		return err
	}

	if res.InsertedID == nil {
		return ErrNoResult
	}

	return nil
}

func (r *MixedProductRepository) Update(id int, p *entity.Product) error {
	complete := r.tryCompleteProduct(p)
	coll := r.noSqlDB.Collection(UserProductsCollection)

	res, err := coll.ReplaceOne(r.ctx, bson.M{"id": id}, complete)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 || res.ModifiedCount == 0 {
		return ErrNoResult
	}

	return nil
}

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
