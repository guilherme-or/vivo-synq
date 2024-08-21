package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/kafka/internal/handler"
)

type AutoOffsetReset string

const (
	AOROldest   AutoOffsetReset = "oldest"
	AOREarliest AutoOffsetReset = "earliest"
	AORNone     AutoOffsetReset = "none"
)

type KafkaConsumer struct {
	*kafka.Consumer
}

func New(server, groupID string, autoOffsetReset AutoOffsetReset) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupID,
		"auto.offset.reset": string(autoOffsetReset),
	})

	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{Consumer: c}, nil
}

func (c *KafkaConsumer) Read(handler *handler.KafkaMessageHandler) {
	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			handler.OnFail(msg, err)
			continue
		}

		handler.OnMessage(msg)
	}
}
