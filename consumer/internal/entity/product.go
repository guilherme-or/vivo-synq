package entity

type Product struct {
	ID               int           `json:"id" bson:"id"`
	Status           string        `json:"status" bson:"status"`
	ProductName      string        `json:"product_name" bson:"product_name"`
	ProductType      string        `json:"product_type" bson:"product_type"`
	SubscriptionType string        `json:"subscription_type" bson:"subscription_type"`
	StartDate        int64         `json:"start_date" bson:"start_date"`
	EndDate          int64         `json:"end_date" bson:"end_date"`
	UserID           *int          `json:"user_id" bson:"user_id"`
	ParentProductID  *int          `json:"parent_product_id" bson:"parent_product_id"`
	SubProducts      []Product     `json:"sub_products" bson:"sub_products"`
	Tags             []string      `json:"tags" bson:"tags"`
	Identifiers      []string      `json:"identifiers" bson:"identifiers"`
	Descriptions     []Description `json:"descriptions" bson:"descriptions"`
	Prices           []Price       `json:"prices" bson:"prices"`
}
