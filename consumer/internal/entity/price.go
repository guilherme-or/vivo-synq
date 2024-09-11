package entity

type Price struct {
	ID              int     `json:"-" bson:"id"`
	ProductID       int     `json:"-" bson:"product_id"`
	Description     string  `json:"description,omitempty" bson:"description"`
	Type            string  `json:"type,omitempty" bson:"type"`
	RecurringPeriod string  `json:"recurring_period,omitempty" bson:"recurring_period"`
	Amount          float64 `json:"amount" bson:"amount"`
}