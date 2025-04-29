package handler

import (
	"context"
	"net/http"

	"github.com/daniil717/adprogass2/api-gateway/client"
	"github.com/daniil717/adprogass2/api-gateway/proto/inventory"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Client *client.InventoryClient
}

func NewProductHandler(c *client.InventoryClient) *ProductHandler {
	return &ProductHandler{Client: c}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		Quantity    int32   `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.Client.Client.CreateProduct(context.Background(), &inventory.ProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"productId": resp.ProductId})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	productId := c.Param("id")

	resp, err := h.Client.Client.GetProduct(context.Background(), &inventory.ProductID{
		Id: productId,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"productId":   resp.ProductId,
		"name":        resp.Name,
		"description": resp.Description,
		"price":       resp.Price,
		"quantity":    resp.Quantity,
	})
}
