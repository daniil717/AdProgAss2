package main

import (
	"context"
	"log"
	"net"
	"os"

	"inventory_service/internal/handler"
	"inventory_service/internal/repo"
	"inventory_service/internal/usecase"
	pb "inventory_service/proto"

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
		log.Fatalf("MongoDB error: %v", err)
	}
	db := client.Database("FoodStore")

	repository := repo.NewMongoRepo(db)
	usecase := usecase.NewProductUseCase(repository)
	handler := handler.NewInventoryHandler(usecase)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, handler)

	log.Println("Inventory Service started on :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
