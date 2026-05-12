package router

import "github.com/gin-gonic/gin"

type OrderItemHandler interface {
	GetAll(c *gin.Context)
}

func OrderItemRounter(rg *gin.RouterGroup, handler OrderItemHandler) {
	orderItems := rg.Group("/orderitems")
	orderItems.GET("/all", handler.GetAll)
}
