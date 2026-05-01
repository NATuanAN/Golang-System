package container

import (
	"go-project/internal/handler"
	"go-project/internal/redis"
	"go-project/internal/repository"
	"go-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
}
type productContainer struct {
	ProductHandler
}

func NewProductContainer(db *gorm.DB, redis redis.Cache) *productContainer {
	repo := repository.NewProductRepo(db)
	service := service.NewProductService(repo)
	handler := handler.NewProductHandler(service, redis)

	return &productContainer{handler}
}
