package handler

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
)

type KafkaMessage struct {
	Payload struct {
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
		Source struct {
			Version   string `json:"version"`
			Connector string `json:"connector"`
			Name      string `json:"name"`
			TimeInMS  int64  `json:"ts_ms"`
			Database  string `json:"db"`
			Schema    string `json:"schema"`
			Table     string `json:"table"`
		} `json:"source"`
		Transaction bool `json:"transaction"`
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

	// TODO: Check message after and before type (Product, Price, Identifier, Description, Tag)

	if message.Payload.Source.Table == "product" {
		fmt.Println("Product")
		// productAfter := message.Payload.After.(*entity.Product)
		// productBefore := message.Payload.Before.(*entity.Product)
	} else if message.Payload.Source.Table == "price" {
		fmt.Println("Price")
	} else if message.Payload.Source.Table == "identifier" {
		fmt.Println("Identifier")
	} else if message.Payload.Source.Table == "description" {
		fmt.Println("Description")
	} else if message.Payload.Source.Table == "tag" {
		fmt.Println("Tag")
	} else {
		fmt.Println("Unknown")
	}

	// afterID := message.Payload.After.ID
	// beforeID := message.Payload.Before.ID

	// if afterID <= 0 && beforeID <= 0 { // Invalid ID value
	// 	fmt.Print("Invalid ID")
	// 	return
	// } else if afterID > 0 && beforeID > 0 { // Update
	// 	fmt.Print("UPDATE")
	// 	if err := h.productRepo.Update(beforeID, &message.Payload.After); err != nil {
	// 		fmt.Printf("...ERROR: %v\n", err)
	// 		return
	// 	}
	// } else if afterID > 0 && beforeID <= 0 { // Insert
	// 	fmt.Print("INSERT")
	// 	if err := h.productRepo.Insert(&message.Payload.After); err != nil {
	// 		fmt.Printf("...ERROR: %v\n", err)
	// 		return
	// 	}
	// } else if afterID <= 0 && beforeID > 0 { // Delete
	// 	fmt.Print("DELETE")
	// 	if err := h.productRepo.Delete(beforeID, message.Payload.Before.ProductType); err != nil {
	// 		fmt.Printf("...ERROR: %v\n", err)
	// 		return
	// 	}
	// }

	// fmt.Printf("...OK! Before: %d After: %d\n", beforeID, afterID)
}

func (h *KafkaMessageHandler) OnFail(msg *kafka.Message, err error) {
	fmt.Printf("FAIL: %v (%v)\n", err, msg)
}
