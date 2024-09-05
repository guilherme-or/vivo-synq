package entity

type Identifier struct {
	ID        int    `json:"-" bson:"id"`
	ProductID int    `json:"-" bson:"product_id"`
	UserID    int    `json:"-" bson:"user_id"`
	Value     string `json:"value" bson:"value"`
}
