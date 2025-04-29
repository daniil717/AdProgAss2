package client

import (
	"log"

	"github.com/daniil717/adprogass2/api-gateway/proto/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	Client user.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	return &UserClient{
		Client: user.NewUserServiceClient(conn),
	}
}
