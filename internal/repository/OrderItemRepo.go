package repository

import (
	"context"
	"errors"
	"go-project/internal/model"

	"gorm.io/gorm"
)

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) *orderItemRepo {
	return &orderItemRepo{db}
}

func (r *orderItemRepo) FindAll(ctx context.Context) ([]model.OrderItem, error) {
	var items []model.OrderItem
	if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return items, nil
}
