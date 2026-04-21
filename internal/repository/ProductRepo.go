package repository

import (
	"context"
	"errors"
	"fmt"
	"go-project/internal/model"
	"go-project/pkg/apperror"

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

func (r *productRepo) FindById(ctx context.Context, id int) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).First(&product, "productid = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, fmt.Errorf("userRepo.FindByID: %w", err)
	}

	return &product, nil
}
