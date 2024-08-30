package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guilherme-or/vivo-synq/consumer/internal/connector"
	"github.com/guilherme-or/vivo-synq/consumer/internal/consumer"
	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/handler"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	// Environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading environment file (.env): " + err.Error())
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
	if err := cs.Subscribe(os.Getenv("KAFKA_TOPIC"), nil); err != nil {
		panic(err)
	}
	defer cs.Close()
	log.Println("Consumer subscribed. Starting to read messages...")

	// Database connection instance
	conn := database.NewMongoDBConn(os.Getenv("NOSQL_URI"))
	// conn := database.NewPostgreSQLConn(os.Getenv("SQL_URI"))
	if err := conn.Open(); err != nil {
		panic(err)
	}
	defer conn.Close()

	// Repository and Handler instance
	repo := repository.NewMongoDBProductRepository(conn.(*database.MongoDBConn))
	// repo := repository.NewPostgreSQLProductRepository(conn.(*database.PostgreSQLConn))
	h := handler.NewKafkaMessageHandler(repo)
	log.Println("Message handler created with database repository...")

	cs.Read(h)
	fmt.Scan()
}
