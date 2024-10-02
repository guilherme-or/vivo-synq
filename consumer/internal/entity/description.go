package entity

type Description struct {
	ID        int    `json:"id" bson:"id"`
	ProductID int    `json:"product_id" bson:"product_id"`
	Text      string `json:"text" bson:"text"`
	URL       string `json:"url,omitempty" bson:"url"`
	Category  string `json:"category,omitempty" bson:"category"`
}
