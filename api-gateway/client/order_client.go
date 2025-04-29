package client

import (
	"log"

	"github.com/daniil717/adprogass2/api-gateway/proto/order"

	"google.golang.org/grpc"
)

type OrderClient struct {
	Client order.OrderServiceClient
}

func NewOrderClient(address string) *OrderClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to order service: %v", err)
	}
	return &OrderClient{
		Client: order.NewOrderServiceClient(conn),
	}
}
