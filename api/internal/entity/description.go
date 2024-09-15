package entity

type Description struct {
	ID        int    `json:"-" bson:"id"`
	ProductID int    `json:"-" bson:"product_id"`
	Text      string `json:"text" bson:"text"`
	URL       string `json:"url,omitempty" bson:"url"`
	Category  string `json:"category,omitempty" bson:"category"`
}

type DescriptionDTO struct {
	Text      string `json:"text" bson:"text"`
	URL       string `json:"url,omitempty" bson:"url"`
	Category  string `json:"category,omitempty" bson:"category"`
}