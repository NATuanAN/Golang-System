package container

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"gorm.io/gorm"
)

type userContainer struct {
	UserHandler handler.UserHandler
}

func NewUserContainer(db *gorm.DB) *userContainer {
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &userContainer{UserHandler: userHandler}
}
