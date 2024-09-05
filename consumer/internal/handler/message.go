package handler

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
)

type KafkaMessage struct {
	Payload struct {
		After  entity.Product `json:"after"`
		Before entity.Product `json:"before"`
	} `json:"payload"`
}

type KafkaMessageHandler struct {
	productRepo repository.ProductRepository
}

func NewKafkaMessageHandler(productRepo repository.ProductRepository) *KafkaMessageHandler {
	return &KafkaMessageHandler{productRepo: productRepo}
}

func (h *KafkaMessageHandler) OnMessage(msg *kafka.Message) {
	var message KafkaMessage

	if err := json.Unmarshal(msg.Value, &message); err != nil {
		fmt.Printf("Error unmarshalling message: %v\n", err)
		return
	}

	// JSON Pretty Print
	jsonBytes, err := json.MarshalIndent(string(msg.Value), "", "    ")
	if err != nil {
		fmt.Printf("Error marshalling message: %v\n", err)
		return
	}
	fmt.Println(string(jsonBytes))

	afterID := message.Payload.After.ID
	beforeID := message.Payload.Before.ID

	if afterID <= 0 && beforeID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if afterID > 0 && beforeID > 0 { // Update
		fmt.Print("UPDATE")
		if err := h.productRepo.Update(beforeID, &message.Payload.After); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if afterID > 0 && beforeID <= 0 { // Insert
		fmt.Print("INSERT")
		if err := h.productRepo.Insert(&message.Payload.After); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if afterID <= 0 && beforeID > 0 { // Delete
		fmt.Print("DELETE")
		if err := h.productRepo.Delete(beforeID, message.Payload.Before.ProductType); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Printf("...OK! Before: %d After: %d\n", beforeID, afterID)
}

func (h *KafkaMessageHandler) OnFail(msg *kafka.Message, err error) {
	fmt.Printf("FAIL: %v (%v)\n", err, msg)
}
