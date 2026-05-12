package repository

import (
	"context"
	"errors"
	"go-project/internal/model"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db}
}

func (o *orderRepo) FindAll(ctx context.Context) ([]model.Order, error) {
	var listOfOrder []model.Order
	if err := o.db.WithContext(ctx).Find(&listOfOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return listOfOrder, nil
}
