package handler

import (
	"context"
	"go-project/internal/model"
	"go-project/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetById(ctx context.Context, id int) (*model.Product, error)
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

func (h *productHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	products, err := h.service.GetById(c, id)
	response.Response(c, products, err)
}
