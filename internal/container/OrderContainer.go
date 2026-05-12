package container

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler interface {
	GetAll(c *gin.Context)
}

type orderContainer struct {
	OrderHandler OrderHandler
}

func NewOrderContainer(db *gorm.DB) *orderContainer {
	orderRepo := repository.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	return &orderContainer{OrderHandler: orderHandler}
}
