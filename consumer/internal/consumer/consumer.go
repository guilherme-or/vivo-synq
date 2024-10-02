package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/consumer/internal/handler"
)

type AutoOffsetReset string

const (
	AORLatest   AutoOffsetReset = "latest"
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

func (c *KafkaConsumer) ReadTimeout(handler *handler.KafkaMessageHandler, t time.Duration) {
	end := time.Now().Add(t)
	for time.Now().Before(end) {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			handler.OnFail(msg, err)
			continue
		}

		handler.OnMessage(msg)
	}
}
