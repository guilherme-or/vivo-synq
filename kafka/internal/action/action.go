package action

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var OnMessage func(*kafka.Message) = func(msg *kafka.Message) {
	fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
}

var OnFail func(*kafka.Message, error) = func(msg *kafka.Message, err error) {
	fmt.Printf("Consumer error: %v (%v)\n", err, msg)
}