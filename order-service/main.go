package main

import (
	"context"
	"log"
	"net"
	"os"

	"order_service/internal/events" // 👈 добавляем events
	"order_service/internal/handler"
	"order_service/internal/logger"
	"order_service/internal/repo"
	"order_service/internal/usecase"
	pb "order_service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {

	logger.Init()

	// MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://mongo:27017" // ⚠️ docker-сервис "mongo"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Mongo connection error: %v", err)
	}
	db := client.Database("FoodStore")

	// NATS
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://nats:4222" // ⚠️ docker-сервис "nats"
	}

	publisher, err := events.NewNatsOrderPublisher(natsURL)
	if err != nil {
		log.Fatalf("NATS connection error: %v", err)
	}

	// Clean Arch layers
	repo := repo.NewMongoRepo(db)
	uc := usecase.NewOrderUseCase(repo, publisher) // 👈 передаём publisher
	h := handler.NewOrderHandler(uc)

	// gRPC
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
