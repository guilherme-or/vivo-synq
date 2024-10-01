package handler

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
)

type (
	TableName    string
	KafkaMessage struct {
		Payload struct {
			After  json.RawMessage `json:"after"`
			Before json.RawMessage `json:"before"`
			Source struct {
				Version      string `json:"version"`
				Connector    string `json:"connector"`
				Name         string `json:"name"`
				TimeInMillis int64  `json:"ts_ms"`
				Database     string `json:"db"`
				Schema       string `json:"schema"`
				Table        string `json:"table"`
			} `json:"source"`
			Transaction bool `json:"transaction"`
		} `json:"payload"`
	}
)

const (
	ProductsTableName     TableName = "products"
	PricesTableName       TableName = "prices"
	IdentifiersTableName  TableName = "identifiers"
	DescriptionsTableName TableName = "descriptions"
	TagsTableName         TableName = "tags"
)

type KafkaMessageHandler struct {
	productRepo     repository.ProductRepository
	priceRepo       repository.PriceRepository
	identifiersRepo repository.IdentifierRepository
	tagsRepo        repository.TagRepository
	descriptionRepo repository.DescriptionRepository
}

func NewKafkaMessageHandler(
	productRepo repository.ProductRepository,
	priceRepo repository.PriceRepository,
	identifiersRepo repository.IdentifierRepository,
	tagsRepo repository.TagRepository,
	descriptionRepo repository.DescriptionRepository,
) *KafkaMessageHandler {
	return &KafkaMessageHandler{
		productRepo:     productRepo,
		priceRepo:       priceRepo,
		identifiersRepo: identifiersRepo,
		tagsRepo:        tagsRepo,
		descriptionRepo: descriptionRepo,
	}
}

func (h *KafkaMessageHandler) OnMessage(msg *kafka.Message) {
	var message KafkaMessage

	if err := json.Unmarshal(msg.Value, &message); err != nil {
		fmt.Printf("Error unmarshalling message: %v\n", err)
		return
	}

	// JSON Pretty Print
	// jsonBytes, err := json.MarshalIndent(string(msg.Value), "", "    ")
	// if err != nil {
	// 	fmt.Printf("Error marshalling message: %v\n", err)
	// 	return
	// }
	// fmt.Println(string(jsonBytes))

	h.processCDC(TableName(message.Payload.Source.Table), message.Payload.After, message.Payload.Before)
}

func (h *KafkaMessageHandler) OnFail(msg *kafka.Message, err error) {
	// TODO: Implement error handling, logging, messaging, etc.
	fmt.Printf("FAIL: %v (%v)\n", err, msg)
}

// Process the CDC (Capture data change) on Products, Prices, Identifiers, Descriptions and Tags
func (h *KafkaMessageHandler) processCDC(table TableName, after, before json.RawMessage) {
	switch table {
	case ProductsTableName:
		var a *entity.Product
		var b *entity.Product
		if err := json.Unmarshal(after, &a); err != nil {
			fmt.Printf("Error unmarshalling %s after: %v\n", ProductsTableName, err)
			return
		}
		if err := json.Unmarshal(before, &b); err != nil {
			fmt.Printf("Error unmarshalling %s before: %v\n", ProductsTableName, err)
			return
		}
		h.cdcProducts(b, a)
	case PricesTableName:
		var a *entity.Price
		var b *entity.Price
		if err := json.Unmarshal(after, &a); err != nil {
			fmt.Printf("Error unmarshalling %s after: %v\n", PricesTableName, err)
			return
		}
		if err := json.Unmarshal(before, &b); err != nil {
			fmt.Printf("Error unmarshalling %s before: %v\n", PricesTableName, err)
			return
		}
		h.cdcPrices(b, a)
	case IdentifiersTableName:
		var a *entity.Identifier
		var b *entity.Identifier
		if err := json.Unmarshal(after, &a); err != nil {
			fmt.Printf("Error unmarshalling %s after: %v\n", IdentifiersTableName, err)
			return
		}
		if err := json.Unmarshal(before, &b); err != nil {
			fmt.Printf("Error unmarshalling %s before: %v\n", IdentifiersTableName, err)
			return
		}
		h.cdcIdentifiers(b, a)
	case DescriptionsTableName:
		var a *entity.Description
		var b *entity.Description
		if err := json.Unmarshal(after, &a); err != nil {
			fmt.Printf("Error unmarshalling %s after: %v\n", DescriptionsTableName, err)
			return
		}
		if err := json.Unmarshal(before, &b); err != nil {
			fmt.Printf("Error unmarshalling %s before: %v\n", DescriptionsTableName, err)
			return
		}
		h.cdcDescriptions(b, a)
	case TagsTableName:
		var a *entity.Tag
		var b *entity.Tag
		if err := json.Unmarshal(after, &a); err != nil {
			fmt.Printf("Error unmarshalling %s after: %v\n", TagsTableName, err)
			return
		}
		if err := json.Unmarshal(before, &b); err != nil {
			fmt.Printf("Error unmarshalling %s before: %v\n", TagsTableName, err)
			return
		}
		h.cdcTags(b, a)
	default:
		fmt.Printf("No handler for table %s\n", table)
	}
}

func (h *KafkaMessageHandler) cdcProducts(before, after *entity.Product) {
	if after == nil && before == nil { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after != nil && before != nil { // Update
		fmt.Printf("UPDATE %s", ProductsTableName)
		if err := h.productRepo.Update(before, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after != nil && before == nil { // Insert
		fmt.Printf("INSERT %s", ProductsTableName)
		if err := h.productRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after == nil && before != nil { // Delete
		fmt.Printf("DELETE %s", ProductsTableName)
		if err := h.productRepo.Delete(before); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Print("...OK!\n")
}

func (h *KafkaMessageHandler) cdcPrices(before, after *entity.Price) {
	if after == nil && before == nil { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after != nil && before != nil { // Update
		fmt.Printf("UPDATE %s", PricesTableName)
		if err := h.priceRepo.Update(before, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after != nil && before == nil { // Insert
		fmt.Printf("INSERT %s", PricesTableName)
		if err := h.priceRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after == nil && before != nil { // Delete
		fmt.Printf("DELETE %s", PricesTableName)
		if err := h.priceRepo.Delete(before); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Print("...OK!\n")
}

func (h *KafkaMessageHandler) cdcIdentifiers(before, after *entity.Identifier) {
	if after == nil && before == nil { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after != nil && before != nil { // Update
		fmt.Printf("UPDATE %s", IdentifiersTableName)
		if err := h.identifiersRepo.Update(before, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after != nil && before == nil { // Insert
		fmt.Printf("INSERT %s", IdentifiersTableName)
		if err := h.identifiersRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after == nil && before != nil { // Delete
		fmt.Printf("DELETE %s", IdentifiersTableName)
		if err := h.identifiersRepo.Delete(before); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Print("...OK!\n")
}

func (h *KafkaMessageHandler) cdcDescriptions(before, after *entity.Description) {
	if after == nil && before == nil { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after != nil && before != nil { // Update
		fmt.Printf("UPDATE %s", DescriptionsTableName)
		if err := h.descriptionRepo.Update(before, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after != nil && before == nil { // Insert
		fmt.Printf("INSERT %s", DescriptionsTableName)
		if err := h.descriptionRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after == nil && before != nil { // Delete
		fmt.Printf("DELETE %s", DescriptionsTableName)
		if err := h.descriptionRepo.Delete(before); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Print("...OK!\n")
}

func (h *KafkaMessageHandler) cdcTags(before, after *entity.Tag) {
	if after == nil && before == nil { // Invalid ID value
		fmt.Print("Invalid ID")
		return
	} else if after != nil && before != nil { // Update
		fmt.Printf("UPDATE %s", TagsTableName)
		if err := h.tagsRepo.Update(before, after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after != nil && before == nil { // Insert
		fmt.Printf("INSERT %s", TagsTableName)
		if err := h.tagsRepo.Insert(after); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	} else if after == nil && before != nil { // Delete
		fmt.Printf("DELETE %s", TagsTableName)
		if err := h.tagsRepo.Delete(before); err != nil {
			fmt.Printf("...ERROR: %v\n", err)
			return
		}
	}

	fmt.Print("...OK!\n")
}
