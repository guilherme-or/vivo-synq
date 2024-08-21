package main

import (
	"fmt"
	"log"

	"github.com/guilherme-or/vivo-synq/kafka/internal/connector"
	"github.com/guilherme-or/vivo-synq/kafka/internal/consumer"
	"github.com/guilherme-or/vivo-synq/kafka/internal/database"
	"github.com/guilherme-or/vivo-synq/kafka/internal/handler"
	"github.com/guilherme-or/vivo-synq/kafka/internal/repository"
)

func main() {
	cs, err := consumer.New("localhost", "myGroup", consumer.AOREarliest)
	if err != nil {
		panic(err)
	}
	log.Println("Consumer created...")

	cn := connector.New("http://localhost:8083/connectors", "vivo_synq-connector", "./debezium-connector.json")
	if err := cn.Register(); err != nil {
		panic(err)
	}
	log.Println("Connector created and registered...")

	if err := cs.Subscribe("postgres.public.products", nil); err != nil {
		panic(err)
	}
	defer cs.Close()
	log.Println("Consumer subscribed. Starting to read messages...")

	conn := database.NewPostgreSQLConn(
		"user=docker password=docker dbname=integrated host=127.0.0.1 port=5433 sslmode=disable",
	)
	if err := conn.Open(); err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.NewPostgreSQLProductRepository(conn.(*database.PostgreSQLConn))
	h := handler.NewKafkaMessageHandler(repo)
	log.Println("Message handler created with database repository...")

	cs.Read(h)
	fmt.Scan()
}
