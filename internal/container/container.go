package container

import (
	"gorm.io/gorm"
)

type Container struct {
	UserContainer *UserContainer
}

func NewContainer(db *gorm.DB) *Container {
	userContainer := NewUserContainer(db)
	return &Container{UserContainer: userContainer}
}
