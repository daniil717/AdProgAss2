package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     string             `bson:"user_id"`
	Items      []OrderItem        `bson:"items"`
	TotalPrice float64            `bson:"total_price"`
	Status     string             `bson:"status"`
}

type OrderItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int32  `bson:"quantity"`
}
