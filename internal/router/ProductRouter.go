package router

import (
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
}

func ProductRounter(rg *gin.RouterGroup, handler ProductHandler) {
	products := rg.Group("/products")
	products.GET("/all", handler.GetAll)
}
