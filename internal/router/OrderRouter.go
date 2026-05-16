package router

import (
	"go-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	GetAll(c *gin.Context)
}

func OrderRounter(rg *gin.RouterGroup, handler OrderHandler) {
	orders := rg.Group("/orders")
	orders.GET("/all", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(), handler.GetAll)
	orders.POST("/create", handler.GetAll)
}
