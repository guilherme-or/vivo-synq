package consumer

import "github.com/confluentinc/confluent-kafka-go/kafka"

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

func (c *KafkaConsumer) Read(onMessage func(*kafka.Message), onFail func(*kafka.Message, error)) {
	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			onFail(msg, err)
			continue
		}

		onMessage(msg)
	}
}
