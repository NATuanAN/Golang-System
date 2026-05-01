package handler

import (
	"context"
	"fmt"
	"go-project/internal/model"
	"go-project/internal/redis"
	"go-project/pkg/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]model.Product, error)
	GetById(ctx context.Context, id int) (*model.Product, error)
}
type productHandler struct {
	service ProductService
	cache   redis.Cache
}

func NewProductHandler(service ProductService, cache redis.Cache) *productHandler {
	return &productHandler{service, cache}
}
func (h *productHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	key := "product:all"

	var products []model.Product

	err := h.cache.Get(ctx, key, &products)
	if err == nil {
		fmt.Println("Cache hit")
		response.Response(c, products, nil)
		return
	}
	if err != redis.ErrCacheMiss {
		fmt.Println("Redis error:", err)
		response.Response(c, nil, err)
		return
	}

	fmt.Println("Query from DB")
	products, err = h.service.GetAll(c)
	if err != nil {
		response.Response(c, nil, err)
		return
	}

	_ = h.cache.Set(ctx, key, products, 15*time.Minute)

	response.Response(c, products, nil)
}

func (h *productHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Response(c, nil, err)
		return
	}
	products, err := h.service.GetById(c, id)
	response.Response(c, products, err)
}
