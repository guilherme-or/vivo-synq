package entity

type Price struct {
	ID              int     `json:"id" bson:"id"`
	ProductID       int     `json:"product_id" bson:"product_id"`
	Description     string  `json:"description" bson:"description"`
	Type            string  `json:"type" bson:"type"`
	RecurringPeriod string  `json:"recurring_period" bson:"recurring_period"`
	Amount          string `json:"amount" bson:"amount"`
}
