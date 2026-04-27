package service

import (
	"context"
	"fmt"
	"go-project/internal/model"
	"go-project/pkg/apperror"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, req *model.User) (*model.User, error)
}
type UserService interface {
	GetById(ctx context.Context, id string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetById(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userService.GetById %s: %w", id, err)
	}
	return user, nil
}

func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("userService.GetAll:  %w", err)
	}
	return users, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	if user, err := s.repo.FindByEmail(ctx, email); err != nil {
		return nil, fmt.Errorf("userService.GetAll:  %w", err)
	} else {
		return user, nil
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	existing, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", apperror.ErrConflict)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", err)
	}
	user.Password = string(hashedPass)

	created, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", err)
	}

	return created, nil
}
