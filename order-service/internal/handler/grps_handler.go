package handler

import (
	"context"
	"order_service/internal/models"
	"order_service/internal/usecase"
	pb "order_service/proto"
)

type OrderHandler struct {
	pb.UnimplementedOrderServiceServer
	uc *usecase.OrderUseCase
}

func NewOrderHandler(uc *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{uc: uc}
}

func (h *OrderHandler) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	var items []models.OrderItem
	for _, i := range req.Items {
		items = append(items, models.OrderItem{
			ProductID: i.ProductId,
			Quantity:  i.Quantity,
		})
	}

	order := &models.Order{
		UserID:     req.UserId,
		Items:      items,
		TotalPrice: float64(req.TotalPrice), // float32 → float64
		Status:     "PENDING",
	}

	id, err := h.uc.PlaceOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:     id,
		Status: order.Status,
	}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.Order, error) {
	order, err := h.uc.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var items []*pb.OrderItem
	for _, i := range order.Items {
		items = append(items, &pb.OrderItem{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
		})
	}

	return &pb.Order{
		Id:         order.ID.Hex(),
		UserId:     order.UserID,
		Items:      items,
		TotalPrice: float32(order.TotalPrice), // float64 → float32
		Status:     order.Status,
	}, nil
}
