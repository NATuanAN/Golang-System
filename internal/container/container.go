package container

import (
	"go-project/internal/redis"

	"gorm.io/gorm"
)

type Container struct {
	*userContainer
	*productContainer
}

func NewContainer(db *gorm.DB, redis redis.Cache) *Container {
	userContainer := NewUserContainer(db)
	productContainer := NewProductContainer(db, redis)
	return &Container{userContainer, productContainer}
}
