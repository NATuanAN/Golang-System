package router

import (
	"go-project/internal/handler"
	"go-project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRounter(rg *gin.RouterGroup, handler handler.UserHandler) {
	users := rg.Group("/users")
	users.GET("/all", middleware.AuthenticationMiddleware(), handler.GetAll)
	users.GET("/:id", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware(), handler.GetById)
	users.GET("/email/:email", handler.GetByEmail)
	users.POST("/create", handler.CreateUser)
	users.POST("/login", handler.Login)
}
