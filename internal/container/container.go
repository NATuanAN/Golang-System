package container

import (
	"gorm.io/gorm"
)

type Container struct {
	*userContainer
	*productContainer
}

func NewContainer(db *gorm.DB) *Container {
	userContainer := NewUserContainer(db)
	productContainer := NewProductContainer(db)
	return &Container{userContainer, productContainer}
}
