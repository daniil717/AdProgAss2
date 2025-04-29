package main

import (
	"github.com/daniil717/adprogass2/api-gateway/client"
	"github.com/daniil717/adprogass2/api-gateway/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gRPC clients
	userClient := client.NewUserClient("user-service:50051")
	orderClient := client.NewOrderClient("order-service:50052")
	inventoryClient := client.NewInventoryClient("inventory-service:50053")

	// Handlers
	userHandler := handler.NewUserHandler(userClient)
	orderHandler := handler.NewOrderHandler(orderClient)
	productHandler := handler.NewProductHandler(inventoryClient)

	// Routes
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)
	r.GET("/users/profile/:id", userHandler.GetProfile)

	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProduct)

	r.Run(":8080")
}
