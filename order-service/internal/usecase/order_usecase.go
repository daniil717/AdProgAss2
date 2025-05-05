package usecase

import (
	"context"
	"log"
	"order_service/internal/events"
	"order_service/internal/models"
)

type OrderRepo interface {
	Create(ctx context.Context, order *models.Order) (string, error)
	GetByID(ctx context.Context, id string) (*models.Order, error)
}

type OrderUseCase struct {
	repo      OrderRepo
	publisher events.OrderPublisher
}

func NewOrderUseCase(r OrderRepo, p events.OrderPublisher) *OrderUseCase {
	return &OrderUseCase{
		repo:      r,
		publisher: p,
	}
}

func (uc *OrderUseCase) PlaceOrder(ctx context.Context, order *models.Order) (string, error) {
	return uc.repo.Create(ctx, order)
}

func (uc *OrderUseCase) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, order *models.Order) error {
	orderID, err := uc.repo.Create(ctx, order)
	if err != nil {
		return err
	}

	productIDs := []string{}
	for _, p := range order.Items {
		productIDs = append(productIDs, p.ProductID)
	}

	err = uc.publisher.PublishOrderCreated(orderID, productIDs)
	if err != nil {
		log.Printf("❌ Failed to publish order.created: %v", err)
	} else {
		log.Printf("✅ Published order.created for order %s", orderID)
	}

	return nil
}
