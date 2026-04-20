package handler

import (
	"context"
	"go-project/internal/model"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
}
type productHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *productHandler {
	return &productHandler{service}
}
func (h *productHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll(c)
	response.Response(c, products, err)
}
