package handler

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/kafka/internal/entity"
	"github.com/guilherme-or/vivo-synq/kafka/internal/repository"
)

type KafkaMessageHandler struct {
	productRepo repository.ProductRepository
}

func NewKafkaMessageHandler(productRepo repository.ProductRepository) *KafkaMessageHandler {
	return &KafkaMessageHandler{productRepo: productRepo}
}

func (h *KafkaMessageHandler) OnMessage(msg *kafka.Message) {
	var message struct {
		Payload struct {
			After  entity.Product `json:"after"`
			Before entity.Product `json:"before"`
		} `json:"payload"`
	}

	if err := json.Unmarshal(msg.Value, &message); err != nil {
		fmt.Printf("Error unmarshalling message: %v\n", err)
		return
	}

	// JSON Pretty Print
	// jsonBytes, err := json.MarshalIndent(message, "", "    ")
	// if err != nil {
	// 	fmt.Printf("Error marshalling message: %v\n", err)
	// 	return
	// }
	// fmt.Println(string(jsonBytes))

	if message.Payload.After.ID <= 0 && message.Payload.Before.ID <= 0 {
		fmt.Println("Product action: ERROR No ID value")
		return
	} else if message.Payload.After.ID > 0 && message.Payload.Before.ID > 0 {
		// Update
		fmt.Println("Product action: UPDATE")
		if err := h.productRepo.Update(message.Payload.Before.ID, &message.Payload.After); err != nil {
			fmt.Printf("Error updating product: %v\n", err)
			return
		}
	} else if message.Payload.After.ID > 0 && message.Payload.Before.ID <= 0 {
		// Insert
		fmt.Println("Product action: INSERT")
		if err := h.productRepo.Insert(&message.Payload.After); err != nil {
			fmt.Printf("Error inserting product: %v\n", err)
			return
		}
	} else if message.Payload.After.ID <= 0 && message.Payload.Before.ID > 0 {
		// Delete
		fmt.Println("Product action: DELETE")
		if err := h.productRepo.Delete(message.Payload.Before.ID, message.Payload.Before.ProductType); err != nil {
			fmt.Printf("Error deleting product: %v\n", err)
			return
		}
	}

	fmt.Println("Product integrated successfully! Before ID: ", message.Payload.Before.ID, "After ID: ", message.Payload.After.ID)
}

func (h *KafkaMessageHandler) OnFail(msg *kafka.Message, err error) {
	fmt.Printf("Consumer error: %v (%v)\n", err, msg)
}
