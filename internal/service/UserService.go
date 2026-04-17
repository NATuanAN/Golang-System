package service

import (
	"context"
	"fmt"
	"go-project/internal/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}
type UserService interface {
	GetById(ctx context.Context, id string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetById(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetById %s: %w", id, err)
	}
	return user, nil
}

func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	user, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetAll have error:  %w", err)
	}
	return user, nil
}
