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
	productRepo     repository.ProductRepository
	priceRepo       repository.PriceRepository
	identifiersRepo repository.IdentifierRepository
	tagsRepo        repository.TagRepository
	descriptionRepo repository.DescriptionRepository
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

	checkTableChanged(message, h)
}

func (h *KafkaMessageHandler) OnFail(msg *kafka.Message, err error) {
	fmt.Printf("FAIL: %v (%v)\n", err, msg)
}

func checkTableChanged(message KafkaMessage, h *KafkaMessageHandler) {

	if message.Payload.Source.Table == "products" {

		after := message.Payload.After.(*entity.Product)
		before := message.Payload.Before.(*entity.Product)

		actionTakerProducts(after, before, h)
	}
	if message.Payload.Source.Table == "prices" {

		after := message.Payload.After.(*entity.Price)
		before := message.Payload.Before.(*entity.Price)

		actionTakerPrices(after, before, h)
	}
	if message.Payload.Source.Table == "identifiers" {

		after := message.Payload.After.(*entity.Identifiers)
		before := message.Payload.Before.(*entity.Identifiers)

		actionTakerIdentifiers(after, before, h)
	}
	if message.Payload.Source.Table == "descriptions" {

		after := message.Payload.After.(*entity.Description)
		before := message.Payload.Before.(*entity.Description)

		actionTakerDescriptions(after, before, h)
	}
	if message.Payload.Source.Table == "tags" {

		after := message.Payload.After.(*entity.Tags)
		before := message.Payload.Before.(*entity.Tags)

		actionTakerTags(after, before, h)
	}
}

func actionTakerProducts(after *entity.Product, before *entity.Product, h *KafkaMessageHandler) {

	if after.ID <= 0 && before.ID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after.ID > 0 && before.ID > 0 { // Update
		fmt.Print("UPDATE")
		if err := h.productRepo.Update(before.ID, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after.ID > 0 && before.ID <= 0 { // Insert
		fmt.Print("INSERT")
		if err := h.productRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after.ID <= 0 && before.ID > 0 { // Delete
		fmt.Print("DELETE")
		if err := h.productRepo.Delete(before.ID, before.ProductType); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}
	fmt.Printf("...OK! Before: %d After: %d\n", before.ID, after.ID)
}

func actionTakerPrices(after *entity.Price, before *entity.Price, h *KafkaMessageHandler) {

	if after.ID <= 0 && before.ID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	}
	fmt.Print("UPDATE")
	if err := h.priceRepo.Update(after); err != nil {
		fmt.Printf("...ERROR: %v\n", err)
		return
	}

	fmt.Printf("...OK! Before: %d After: %d\n", before.ID, after.ID)
}

func actionTakerIdentifiers(after *entity.Identifiers, before *entity.Identifiers, h *KafkaMessageHandler) {

	if after.ID <= 0 && before.ID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	}
	if err := h.identifiersRepo.Update(after); err != nil {
		fmt.Printf("...ERROR: %v\n", err)
		return
	}

	fmt.Printf("...OK! Before: %d After: %d\n", before.ID, after.ID)
}

func actionTakerDescriptions(after *entity.Description, before *entity.Description, h *KafkaMessageHandler) {

	if after.ID <= 0 && before.ID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	}
	fmt.Print("UPDATE")
	if err := h.descriptionRepo.Update(after); err != nil {
		fmt.Printf("...ERROR: %v\n", err)
		return
	}
	fmt.Printf("...OK! Before: %d After: %d\n", before.ID, after.ID)
}

func actionTakerTags(after *entity.Tags, before *entity.Tags, h *KafkaMessageHandler) {

	if after.ID <= 0 && before.ID <= 0 { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	}
	fmt.Print("UPDATE")
	if err := h.tagsRepo.Update(after); err != nil {
		fmt.Printf("...ERROR: %v\n", err)
		return
	}

	fmt.Printf("...OK! Before: %d After: %d\n", before.ID, after.ID)
}
