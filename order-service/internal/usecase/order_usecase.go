package usecase

import (
	"context"
	"order_service/internal/models"
)

type OrderRepo interface {
	Create(ctx context.Context, order *models.Order) (string, error)
	GetByID(ctx context.Context, id string) (*models.Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
}

func NewOrderUseCase(r OrderRepo) *OrderUseCase {
	return &OrderUseCase{repo: r}
}

func (u *OrderUseCase) PlaceOrder(ctx context.Context, order *models.Order) (string, error) {
	order.Status = "PENDING"
	return u.repo.Create(ctx, order)
}

func (u *OrderUseCase) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return u.repo.GetByID(ctx, id)
}
