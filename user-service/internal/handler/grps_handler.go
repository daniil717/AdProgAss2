package handler

import (
	"context"
	"user_service/internal/models"
	"user_service/internal/usecase"
	pb "user_service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(), 
	}
	err := h.uc.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: user.ID, Message: "User registered"}, nil
}

func (h *UserHandler) AuthenticateUser(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	ok, err := h.uc.Authenticate(ctx, req.GetEmail(), req.GetPassword())
	return &pb.AuthResponse{Success: ok, Token: "dummy-token"}, err
}
