package repo

import (
	"context"
	"order_service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) *MongoRepo {
	return &MongoRepo{col: db.Collection("orders")}
}

func (r *MongoRepo) Create(ctx context.Context, order *models.Order) (string, error) {
	res, err := r.col.InsertOne(ctx, order)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *MongoRepo) GetByID(ctx context.Context, id string) (*models.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order models.Order
	err = r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	return &order, err
}
