package repo

import (
	"context"
	"inventory_service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(db *mongo.Database) *MongoRepo {
	return &MongoRepo{col: db.Collection("products")}
}

func (r *MongoRepo) Create(ctx context.Context, product *models.Product) error {
	_, err := r.col.InsertOne(ctx, product)
	return err
}

func (r *MongoRepo) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.col.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	return &product, err
}

func (r *MongoRepo) GetAll(ctx context.Context) ([]*models.Product, error) {
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []*models.Product
	for cur.Next(ctx) {
		var p models.Product
		if err := cur.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *MongoRepo) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.col.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
