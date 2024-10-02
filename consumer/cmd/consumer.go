package main

import (
	"log"
	"os"
	"strings"

	"github.com/guilherme-or/vivo-synq/consumer/internal/connector"
	"github.com/guilherme-or/vivo-synq/consumer/internal/consumer"
	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/handler"
	"github.com/guilherme-or/vivo-synq/consumer/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	// Environment variables
	if os.Getenv("CONNECTOR_URL") == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	// Consumer instance
	cs, err := consumer.New(os.Getenv("CONSUMER_HOST"), os.Getenv("CONSUMER_GROUP_ID"), consumer.AOREarliest)
	if err != nil {
		panic(err)
	}

	log.Println("Consumer created...")

	// Connector instance
	cn := connector.New(os.Getenv("CONNECTOR_URL"), os.Getenv("CONNECTOR_NAME"), os.Getenv("CONNECTOR_FILE"))
	if err := cn.Register(); err != nil {
		panic(err)
	}
	log.Println("Connector created and registered...")

	// Consumer subscription to topic
	topics := strings.Split(os.Getenv("KAFKA_TOPICS"), ",")
	if err := cs.SubscribeTopics(topics, nil); err != nil {
		panic(err)
	}
	defer cs.Close()
	log.Println("Consumer subscribed. Starting to read messages...")

	// Database connection instance
	// sqlConn, noSqlConn := databaseConnection()
	noSqlConn := databaseConnection()
	// defer sqlConn.Close()
	defer noSqlConn.Close()

	// Repository and Handler instance
	h := handlerInstance(noSqlConn)
	log.Println("Message handler created with database repository...")

	// Infinite loop to read messages
	cs.Read(h)
}

func databaseConnection() database.NoSQLConn {
	// sqlConn := database.NewPostgreSQLConn(os.Getenv("SQL_URI"))
	// if err := sqlConn.Open(); err != nil {
	// 	panic(err)
	// }

	noSqlConn := database.NewMongoDBConn(os.Getenv("NOSQL_URI"))
	if err := noSqlConn.Open(); err != nil {
		panic(err)
	}

	// return sqlConn, noSqlConn
	return noSqlConn
}

func handlerInstance(noSqlConn database.NoSQLConn) *handler.KafkaMessageHandler {
	// repo := repository.NewPostgreSQLProductRepository(conn.(*database.PostgreSQLConn))
	// repo := repository.NewMixedProductRepository(sqlConn.(*database.PostgreSQLConn), noSqlConn.(*database.MongoDBConn))
	productRepo := infra.NewMongoDBProductRepository(noSqlConn.(*database.MongoDBConn))
	priceRepo := infra.NewMongoDBPriceRepository(noSqlConn.(*database.MongoDBConn))
	identifierRepo := infra.NewMongoDBIdentifierRepository(noSqlConn.(*database.MongoDBConn))
	tagRepo := infra.NewMongoDBTagRepository(noSqlConn.(*database.MongoDBConn))
	descriptionRepo := infra.NewMongoDBDescriptionRepository(noSqlConn.(*database.MongoDBConn))

	return handler.NewKafkaMessageHandler(
		productRepo,
		priceRepo,
		identifierRepo,
		tagRepo,
		descriptionRepo,
	)
}
