package usecase

import (
	"context"
	"inventory_service/internal/models"
)

type ProductRepo interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	GetAll(ctx context.Context) ([]*models.Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (u *ProductUseCase) Create(ctx context.Context, p *models.Product) error {
	return u.repo.Create(ctx, p)
}

func (u *ProductUseCase) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *ProductUseCase) GetAll(ctx context.Context) ([]*models.Product, error) {
	return u.repo.GetAll(ctx)
}

func (u *ProductUseCase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
