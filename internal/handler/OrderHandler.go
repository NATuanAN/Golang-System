package handler

import (
	"context"
	"go-project/internal/model"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrderService interface {
	GetAll(ctx context.Context) ([]model.Order, error)
}

type orderHandler struct {
	service OrderService
}

func NewOrderHandler(service OrderService) *orderHandler {
	return &orderHandler{service}
}

func (h *orderHandler) GetAll(c *gin.Context) {
	orders, err := h.service.GetAll(c.Request.Context())
	response.Response(c, orders, err)
}
