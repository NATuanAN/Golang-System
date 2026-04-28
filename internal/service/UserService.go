package service

import (
	"context"
	"errors"
	"fmt"
	"go-project/internal/jwt"
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
	Login(ctx context.Context, email string, password string) (string, error)
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
	_, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		if !errors.Is(err, apperror.ErrNotFound) {
			return nil, fmt.Errorf("userService.CreateUser: %w", err)
		}
	} else {
		return nil, apperror.ErrConflict.WithMessage(
			fmt.Sprintf("email %s already exists", user.Email),
		)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser hash password: %w", err)
	}
	user.Password = string(hashedPass)
	created, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("userService.CreateUser: %w", err)
	}

	return created, nil
}

func (s *userService) Login(ctx context.Context, email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", apperror.ErrBadRequest.WithMessage("Email or password must not be null or empty")
	}

	existing, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return "", apperror.ErrUnauthorized.WithMessage("Invalid email or password")
		}
		return "", fmt.Errorf("userService.Login: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(password)); err != nil {
		return "", apperror.ErrUnauthorized.WithMessage("Invalid email or password")
	}

	token, err := jwt.Generate(existing.ID)
	if err != nil {
		return "", fmt.Errorf("userService.Login: %w", err)
	}

	return token, nil
}
