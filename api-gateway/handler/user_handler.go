package handler

import (
	"context"
	"net/http"

	"github.com/daniil717/adprogass2/api-gateway/client"
	"github.com/daniil717/adprogass2/api-gateway/proto/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Client *client.UserClient
}

func NewUserHandler(c *client.UserClient) *UserHandler {
	return &UserHandler{Client: c}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.Client.Client.RegisterUser(context.Background(), &user.UserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": resp.UserId})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := h.Client.Client.AuthenticateUser(context.Background(), &user.AuthRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": resp.Token})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.Param("id")

	resp, err := h.Client.Client.GetUserProfile(context.Background(), &user.UserID{
		Id: userID,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  resp.Name,
		"email": resp.Email,
	})
}
