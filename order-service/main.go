package main

import (
	"context"
	"log"
	"net"
	"os"

	"order_service/internal/handler"
	"order_service/internal/repo"
	"order_service/internal/usecase"
	pb "order_service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Mongo connection error: %v", err)
	}
	db := client.Database("FoodStore")

	repo := repo.NewMongoRepo(db)
	uc := usecase.NewOrderUseCase(repo)
	h := handler.NewOrderHandler(uc)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, h)

	log.Println("Order Service running on :50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}
