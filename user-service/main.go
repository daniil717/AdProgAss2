package main

import (
	"context"
	"log"
	"net"

	"user_service/internal/handler"
	"user_service/internal/repo"
	"user_service/internal/usecase"
	pb "user_service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	mongoURI := "mongodb://mongo:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("MongoDB error: %v", err)
	}
	db := client.Database("FoodStore")

	repository := repo.NewMongoRepo(db)
	usecase := usecase.NewUserUseCase(repository)
	handler := handler.NewUserHandler(usecase)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)

	log.Println("User Service started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
