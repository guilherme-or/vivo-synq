package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/guilherme-or/vivo-synq/gateway/internal/database"
	"github.com/guilherme-or/vivo-synq/gateway/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tolerância de 15 minutos na memória. Valor imutável
// considerado como aceitável pela empresa parceira (Vivo)
const (
	KeyExpTime        = time.Minute * 15
	mongoDatabaseName = "vivo-synq"
	mongoCollName     = "user_products"
)

type MixedProductRepository struct {
	redisClient   *redis.Client
	mongoDBClient *mongo.Client
}

func NewMixedProductRepository(mongoDBConn *database.MongoDBConn, redisConn *database.RedisConn) ProductRepository {
	return &MixedProductRepository{
		redisClient:   redisConn.GetClient().(*redis.Client),
		mongoDBClient: mongoDBConn.GetClient().(*mongo.Client),
	}
}

func (r *MixedProductRepository) getMongoDBUserProductsColl() *mongo.Collection {
	db := r.mongoDBClient.Database(mongoDatabaseName)
	return db.Collection(mongoCollName)
}

func (r *MixedProductRepository) Find(userID string) ([]entity.Product, error) {
	var filter bson.D
	filterKey := "user_id"
	ctx := context.TODO()

	intID, err := strconv.Atoi(userID)
	if err != nil {
		filter = bson.D{{Key: filterKey, Value: userID}}
	} else {
		filter = bson.D{{Key: filterKey, Value: intID}}
	}

	cursor, err := r.getMongoDBUserProductsColl().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []entity.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errors.New("empty products array")
	}

	return products, nil
}

func (r *MixedProductRepository) FindInCache(userID string) ([]entity.Product, error) {
	result := r.redisClient.Get(userID)
	if err := result.Err(); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("no data found in cache")
	}

	var data []byte
	if err := result.Scan(&data); err != nil {
		return nil, err
	}

	var products []entity.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MixedProductRepository) SaveInCache(userID string, products []entity.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	if err := r.redisClient.Set(userID, data, KeyExpTime).Err(); err != nil {
		return err
	}

	return nil
}
