package service

import (
	"context"
	"fmt"
	"go-project/internal/model"
)

type OrderItemRepo interface {
	FindAll(ctx context.Context) ([]model.OrderItem, error)
}

type orderItemService struct {
	repo OrderItemRepo
}

func NewOrderItemService(repo OrderItemRepo) *orderItemService {
	return &orderItemService{repo}
}

func (s *orderItemService) GetAll(ctx context.Context) ([]model.OrderItem, error) {
	items, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("OrderItemService: %w", err)
	}
	return items, nil
}
