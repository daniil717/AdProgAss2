package handler

import (
	"context"
	"inventory_service/internal/models"
	"inventory_service/internal/usecase"
	pb "inventory_service/proto"
)

type InventoryHandler struct {
	pb.UnimplementedInventoryServiceServer
	uc *usecase.ProductUseCase
}

func NewInventoryHandler(uc *usecase.ProductUseCase) *InventoryHandler {
	return &InventoryHandler{uc: uc}
}

func (h *InventoryHandler) CreateProduct(ctx context.Context, req *pb.Product) (*pb.ProductResponse, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	err := h.uc.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{Message: "Product created"}, nil
}

func (h *InventoryHandler) GetProductByID(ctx context.Context, req *pb.ProductID) (*pb.Product, error) {
	product, err := h.uc.GetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Product{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

func (h *InventoryHandler) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductList, error) {
	products, err := h.uc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, p := range products {
		pbProducts = append(pbProducts, &pb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
		})
	}

	return &pb.ProductList{Products: pbProducts}, nil
}

func (h *InventoryHandler) DeleteProduct(ctx context.Context, req *pb.ProductID) (*pb.ProductResponse, error) {
	err := h.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{Message: "Deleted"}, nil
}
