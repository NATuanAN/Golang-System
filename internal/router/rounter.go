package router

import (
	"go-project/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupRouter(container *container.Container) *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())

	v1 := route.Group("/api/v1")
	UserRounter(v1, container.UserHandler)
	ProductRounter(v1, container.ProductHandler)
	OrderRounter(v1, container.OrderHandler)
	OrderItemRounter(v1, container.OrderItemHandler)
	return route
}
