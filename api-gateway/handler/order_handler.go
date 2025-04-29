package handler

import (
	"context"
	"net/http"

	"github.com/daniil717/adprogass2/api-gateway/client"
	"github.com/daniil717/adprogass2/api-gateway/proto/order"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Client *client.OrderClient
}

func NewOrderHandler(c *client.OrderClient) *OrderHandler {
	return &OrderHandler{Client: c}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		UserId    string  `json:"user_id"`
		ProductId string  `json:"product_id"`
		Quantity  int32   `json:"quantity"`
		Price     float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.Client.Client.CreateOrder(context.Background(), &order.OrderRequest{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
		Price:     req.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orderId": resp.OrderId})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderId := c.Param("id")

	resp, err := h.Client.Client.GetOrder(context.Background(), &order.OrderID{
		Id: orderId,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orderId":   resp.OrderId,
		"userId":    resp.UserId,
		"productId": resp.ProductId,
		"quantity":  resp.Quantity,
		"price":     resp.Price,
		"status":    resp.Status,
	})
}
