package main

import (
	"fmt"

	"github.com/guilherme-or/vivo-synq/kafka/internal/action"
	"github.com/guilherme-or/vivo-synq/kafka/internal/connector"
	"github.com/guilherme-or/vivo-synq/kafka/internal/consumer"
)

func main() {
	cs, err := consumer.New("localhost:9092", "1", consumer.AOREarliest)
	if err != nil {
		panic(err)
	}

	cn := connector.New("http://localhost:9093/connectors", "vivo_synq-connector", "./debezium-connector.json")
	if err := cn.Register(); err != nil {
		panic(err)
	}

	if err := cs.Subscribe("postgres.public.products", nil); err != nil {
		panic(err)
	}
	defer cs.Close()

	cs.Read(action.OnMessage, action.OnFail)

	fmt.Scan()
}
