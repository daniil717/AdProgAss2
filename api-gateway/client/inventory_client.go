package client

import (
	"log"

	"github.com/daniil717/adprogass2/api-gateway/proto/inventory"

	"google.golang.org/grpc"
)

type InventoryClient struct {
	Client inventory.InventoryServiceClient
}

func NewInventoryClient(address string) *InventoryClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to inventory service: %v", err)
	}
	return &InventoryClient{
		Client: inventory.NewInventoryServiceClient(conn),
	}
}
