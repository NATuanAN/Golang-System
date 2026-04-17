package container

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"gorm.io/gorm"
)

type UserContainer struct {
	UserHandler handler.UserHandler
}

func NewUserContainer(db *gorm.DB) *UserContainer {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	return &UserContainer{UserHandler: userHandler}
}
