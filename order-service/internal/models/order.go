package models

type OrderItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int32  `bson:"quantity"`
}

type Order struct {
	ID         string      `bson:"_id,omitempty"`
	UserID     string      `bson:"user_id"`
	Items      []OrderItem `bson:"items"`
	TotalPrice float32     `bson:"total_price"`
	Status     string      `bson:"status"`
}
