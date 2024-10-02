package entity

type Identifier struct {
	ID         int    `json:"id" bson:"id"`
	UserID     int    `json:"user_id" bson:"user_id"`
	ProductID  int    `json:"product_id" bson:"product_id"`
	Identifier string `json:"identifier" bson:"identifier"`
}
