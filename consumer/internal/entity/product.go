package entity

type Product struct {
	ID               int            `json:"id" bson:"id"`
	Status           string         `json:"status" bson:"status"`
	ProductName      string         `json:"product_name" bson:"product_name"`
	ProductType      string         `json:"product_type" bson:"product_type"`
	SubscriptionType string         `json:"subscription_type" bson:"subscription_type"`
	StartDate        int64          `json:"start_date" bson:"start_date"`
	EndDate          int64          `json:"end_date" bson:"end_date"`
	UserID           int            `json:"user_id" bson:"user_id"`
	ParentProductID  *int           `json:"parent_product_id,omitempty" bson:"parent_product_id"`
	Tags             *[]string      `json:"tags,omitempty" bson:"tags"`
	Identifiers      *[]string      `json:"identifiers,omitempty" bson:"identifiers"`
	Descriptions     *[]Description `json:"descriptions,omitempty" bson:"descriptions"`
	Prices           *[]Price       `json:"prices,omitempty" bson:"prices"`
}
