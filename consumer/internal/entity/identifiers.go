package entity

type Identifiers struct {
	ID         int    `json:"id" bson:"id"`
	UserID     int    `json:"user_id,omitempty" bson:"user_id"`
	ProductID  int    `json:"product_id" bson:"product_id"`
	Identifier string `json:"identifier" bson:"identifier"`
}
