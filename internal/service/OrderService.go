package service

import (
	"context"
	"fmt"
	"go-project/internal/model"
)

type OrderRepo interface {
	FindAll(ctx context.Context) ([]model.Order, error)
}

type orderService struct {
	repo OrderRepo
}

func NewOrderService(repo OrderRepo) *orderService {
	return &orderService{repo}
}
func (s *orderService) GetAll(ctx context.Context) ([]model.Order, error) {
	orders, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("OrderService: %w", err)
	}
	return orders, nil
}
