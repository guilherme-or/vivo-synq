package entity

type Product struct {
	ID               int  `json:"id"`
	Status           string `json:"status"`
	ProductName      string `json:"product_name"`
	ProductType      string `json:"product_type"`
	SubscriptionType string `json:"subscription_type"`
	StartDate        int64  `json:"start_date"`
	EndDate          int64  `json:"end_date"`
	ParentProductID  int  `json:"parent_product_id"`
	// TODO: Add ManyToOne fields
}