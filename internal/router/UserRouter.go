package router

import (
	"go-project/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRounter(rg *gin.RouterGroup, handler handler.UserHandler) {
	users := rg.Group("/users")
	users.GET("/:id", handler.GetById)
	users.GET("/all", handler.GetAll)
	users.GET("/email/:email", handler.GetByEmail)
	users.POST("/create", handler.CreateUser)
}
