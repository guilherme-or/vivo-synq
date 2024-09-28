package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/guilherme-or/vivo-synq/consumer/internal/connector"
	"github.com/guilherme-or/vivo-synq/consumer/internal/consumer"
	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/handler"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
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
	sqlConn, noSqlConn := databaseConnection()
	defer sqlConn.Close()
	defer noSqlConn.Close()

	// Repository and Handler instance
	// repo := repository.NewMongoDBProductRepository(conn.(*database.MongoDBConn))
	// repo := repository.NewPostgreSQLProductRepository(conn.(*database.PostgreSQLConn))
	repo := repository.NewMixedProductRepository(sqlConn.(*database.PostgreSQLConn), noSqlConn.(*database.MongoDBConn))
	h := handler.NewKafkaMessageHandler(repo)
	log.Println("Message handler created with database repository...")

	cs.Read(h)
	fmt.Scan()
}

func databaseConnection() (database.SQLConn, database.NoSQLConn) {
	// Database connection instance
	sqlConn := database.NewPostgreSQLConn(os.Getenv("SQL_URI"))
	if err := sqlConn.Open(); err != nil {
		panic(err)
	}

	noSqlConn := database.NewMongoDBConn(os.Getenv("NOSQL_URI"))
	if err := noSqlConn.Open(); err != nil {
		panic(err)
	}

	return sqlConn, noSqlConn
}
