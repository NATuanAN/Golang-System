package router

import (
	"go-project/internal/handler"
	"go-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRounter(rg *gin.RouterGroup, handler handler.UserHandler) {
	users := rg.Group("/users")
	users.GET("/:id", handler.GetById)
	users.GET("/all", middleware.AuthMiddleware(), handler.GetAll)
	users.GET("/email/:email", handler.GetByEmail)
	users.POST("/create", handler.CreateUser)
	users.POST("/login", handler.Login)
}
