package repo

import (
	"context"
	"user-service/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) *MongoRepo {
	return &MongoRepo{
		collection: db.Collection("users"),
	}
}

func (r *MongoRepo) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *MongoRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, map[string]interface{}{"email": email}).Decode(&user)
	return &user, err
}
