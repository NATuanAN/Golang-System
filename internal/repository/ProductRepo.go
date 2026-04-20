package repository

import (
	"context"
	"go-project/internal/model"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db}
}

func (r *productRepo) FindAll(ctx context.Context) ([]model.Product, error) {
	var productList []model.Product
	err := r.db.WithContext(ctx).Find(&productList).Error
	if err != nil {
		return nil, err
	}
	return productList, nil
}
