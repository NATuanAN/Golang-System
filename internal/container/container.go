package container

import (
	"go-project/internal/redis"

	"gorm.io/gorm"
)

type Container struct {
	*userContainer
	*productContainer
	*orderContainer
	*orderItemContainer
}

func NewContainer(db *gorm.DB, redis redis.Cache) *Container {
	userContainer := NewUserContainer(db)
	productContainer := NewProductContainer(db, redis)
	orderContainer := NewOrderContainer(db)
	orderItemContainer := NewOrderItemContainer(db)
	return &Container{userContainer, productContainer, orderContainer, orderItemContainer}
}
