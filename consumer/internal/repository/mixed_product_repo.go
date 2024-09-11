package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// Query result example
	// {
	// 	"id":30,
	// 	"status":"activating",
	// 	"product_name":"IPTV Advanced",
	// 	"product_type":"iptv",
	// 	"subscription_type":"postpaid",
	// 	"start_date":1730494800,
	// 	"end_date":null,
	// 	"user_id":7,
	// 	"parent_product_id":null,
	// 	"identifiers":[
	// 	   "ID-012345",
	// 	   "INT-500-555"
	// 	],
	// 	"descriptions":[
	// 	   {
	// 		  "text":"Landline Economy Plan with limited calls",
	// 		  "url":"https://example.com/landline-economy",
	// 		  "category":"dates"
	// 	   }
	// 	],
	// 	"prices":[
	// 	   {
	// 		  "description":"Landline plan upgrade",
	// 		  "type":"one-off",
	// 		  "recurring_period":null,
	// 		  "amount":9.99
	// 	   }
	// 	],
	// 	"sub_products":[
	// 	   {
	// 		  "id":58,
	// 		  "status":"cancelled",
	// 		  "product_name":"IPTV Advanced",
	// 		  "product_type":"iptv",
	// 		  "subscription_type":"postpaid",
	// 		  "start_date":1701432000,
	// 		  "end_date":1704110400,
	// 		  "user_id":11,
	// 		  "parent_product_id":30,
	// 		  "identifiers":null,
	// 		  "descriptions":null,
	// 		  "prices":null
	// 	   }
	// 	]
	// }
	queryGetCompleteProduct = `
		SELECT * FROM get_complete_product($1)
	`
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

	// coll := noSqlDB.Collection(UserProductsCollection)
	// coll.DeleteMany(context.Background(), bson.M{})

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

// MongoDBRepository feature
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

// MongoDBRepository feature
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
