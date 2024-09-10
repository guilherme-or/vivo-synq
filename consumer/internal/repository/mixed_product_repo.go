package repository

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	queryGetProductInformation = `
		SELECT 
		    p.id, p.status, p.product_name, p.product_type, p.subscription_type, 
		    EXTRACT(EPOCH FROM p.start_date) AS start_date, 
		    EXTRACT(EPOCH FROM p.end_date) AS end_date,
		    p.user_id, p.parent_product_id, 
		    array_agg(DISTINCT i.identifier) AS identifiers,
		    array_agg(DISTINCT d.text || '|' || d.url || '|' || d.category) AS descriptions,
		    array_agg(DISTINCT pr.description || '|' || pr.type || '|' || pr.recurring_period || '|' || pr.amount) AS prices
		FROM products p
		LEFT JOIN identifiers i ON p.id = i.product_id
		LEFT JOIN descriptions d ON p.id = d.product_id
		LEFT JOIN prices pr ON p.id = pr.product_id
		WHERE p.id = $1
		GROUP BY p.id;
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

	return &MixedProductRepository{
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
		ctx:     context.TODO(),
	}
}

// TODO: Fix wrong string parsing to product information
func (r *MixedProductRepository) getCompleteProduct(incomplete *entity.Product) *entity.Product {
	// Executando a query
	row := r.sqlDB.QueryRow(queryGetProductInformation, incomplete.ID)

	// Variáveis para os resultados da query
	var product entity.Product
	var startDate, endDate sql.NullFloat64
	var identifiers, descriptions, prices sql.NullString

	// Scan dos resultados
	err := row.Scan(
		&product.ID,
		&product.Status,
		&product.ProductName,
		&product.ProductType,
		&product.SubscriptionType,
		&startDate,
		&endDate,
		&product.UserID,
		&product.ParentProductID,
		&identifiers,
		&descriptions,
		&prices,
	)
	if err != nil {
		return incomplete
	}

	// Processando Identifiers
	if identifiers.Valid {
		idList := strings.Split(identifiers.String, ",")
		product.Identifiers = &idList
	}

	// Processando Descriptions
	if descriptions.Valid {
		var descList []entity.Description
		descEntries := strings.Split(descriptions.String, ",")
		for _, entry := range descEntries {
			parts := strings.Split(entry, "|")
			if len(parts) == 3 {
				descList = append(descList, entity.Description{
					Text:     parts[0],
					URL:      parts[1],
					Category: parts[2],
				})
			}
		}
		product.Descriptions = &descList
	}

	// Processando Prices
	if prices.Valid {
		var priceList []entity.Price
		priceEntries := strings.Split(prices.String, ",")
		for _, entry := range priceEntries {
			parts := strings.Split(entry, "|")
			if len(parts) == 4 {
				amount, _ := strconv.ParseFloat(parts[3], 64)
				priceList = append(priceList, entity.Price{
					Description:     parts[0],
					Type:            parts[1],
					RecurringPeriod: parts[2],
					Amount:          amount,
				})
			}
		}
		product.Prices = &priceList
	}

	return &product
}

// MongoDBRepository feature
func (r *MixedProductRepository) Insert(p *entity.Product) error {
	complete := r.getCompleteProduct(p)
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
	complete := r.getCompleteProduct(p)
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
