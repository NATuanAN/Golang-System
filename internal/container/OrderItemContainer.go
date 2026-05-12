package container

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderItemHandler interface {
	GetAll(c *gin.Context)
}

type orderItemContainer struct {
	OrderItemHandler OrderItemHandler
}

func NewOrderItemContainer(db *gorm.DB) *orderItemContainer {
	repo := repository.NewOrderItemRepo(db)
	service := service.NewOrderItemService(repo)
	handler := handler.NewOrderItemHandler(service)

	return &orderItemContainer{OrderItemHandler: handler}
}
