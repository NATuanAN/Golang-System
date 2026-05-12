package handler

import (
	"context"
	"go-project/internal/model"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrderItemService interface {
	GetAll(ctx context.Context) ([]model.OrderItem, error)
}

type orderItemHandler struct {
	service OrderItemService
}

func NewOrderItemHandler(service OrderItemService) *orderItemHandler {
	return &orderItemHandler{service}
}

func (h *orderItemHandler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll(c.Request.Context())
	response.Response(c, items, err)
}
