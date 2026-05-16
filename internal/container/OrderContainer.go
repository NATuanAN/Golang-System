package container

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"gorm.io/gorm"
)

type orderContainer struct {
	OrderHandler handler.OrderHandler
}

func NewOrderContainer(db *gorm.DB) *orderContainer {
	orderRepo := repository.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	return &orderContainer{OrderHandler: orderHandler}
}
