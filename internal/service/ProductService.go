package service

import (
	"context"
	"fmt"
	"go-project/internal/model"
)

type ProductRepo interface {
	FindAll(ctx context.Context) ([]model.Product, error)
}

type productService struct {
	repo ProductRepo
}

func ProductService(repo ProductRepo) *productService {
	return &productService{repo}
}
func (s *productService) GetAll(ctx context.Context) ([]model.Product, error) {
	productList, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAll have error:  %w", err)
	}
	return productList, nil
}
