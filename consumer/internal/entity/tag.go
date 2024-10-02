package entity

type Tag struct {
	ID        int    `json:"id" bson:"id"`
	ProductId int    `json:"product_id" bson:"product_id"`
	Tag       string `json:"tag" bson:"tag"`
}
