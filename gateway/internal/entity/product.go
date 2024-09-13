package entity

import "time"

// Product representa um produto contratado pelo usuário
// Esse produto é preparado para a resposta da API
type Product struct {
	ID               int            `json:"id" bson:"id"`
	Status           string         `json:"status" bson:"status"`
	ProductName      string         `json:"product_name" bson:"product_name"`
	ProductType      string         `json:"product_type" bson:"product_type"`
	SubscriptionType string         `json:"subscription_type" bson:"subscription_type"`
	StartDate        time.Time      `json:"start_date" bson:"start_date"`
	EndDate          time.Time      `json:"end_date,omitempty" bson:"end_date"`
	UserID           int            `json:"-" bson:"user_id"`
	ParentProductID  *int           `json:"-" bson:"parent_product_id"`
	SubProducts      *[]Product     `json:"sub_products,omitempty" bson:"sub_products"`
	Tags             *[]string      `json:"tags,omitempty" bson:"tags"`
	Identifiers      *[]string      `json:"identifiers,omitempty" bson:"identifiers"`
	Descriptions     *[]Description `json:"descriptions,omitempty" bson:"descriptions"`
	Prices           *[]Price       `json:"prices,omitempty" bson:"prices"`
}

type ProductDTO struct {
	ID               int               `json:"id" bson:"id"`
	Status           string            `json:"status" bson:"status"`
	ProductName      string            `json:"product_name" bson:"product_name"`
	ProductType      string            `json:"product_type" bson:"product_type"`
	SubscriptionType string            `json:"subscription_type" bson:"subscription_type"`
	StartDate        int64             `json:"start_date" bson:"start_date"`
	EndDate          int64             `json:"end_date" bson:"end_date"`
	SubProducts      *[]ProductDTO     `json:"sub_products,omitempty" bson:"sub_products"`
	Tags             *[]string         `json:"tags,omitempty" bson:"tags"`
	Identifiers      *[]string         `json:"identifiers,omitempty" bson:"identifiers"`
	Descriptions     *[]DescriptionDTO `json:"descriptions,omitempty" bson:"descriptions"`
	Prices           *[]PriceDTO       `json:"prices,omitempty" bson:"prices"`
}
