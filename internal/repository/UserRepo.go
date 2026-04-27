package repository

import (
	"context"
	"errors"
	"fmt"
	"go-project/internal/model"
	"go-project/pkg/apperror"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, "userid = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, fmt.Errorf("userRepo.FindByID: %w", err)
	}
	return &user, nil
}

func (r *userRepo) FindAll(ctx context.Context) ([]model.User, error) {
	var user []model.User
	err := r.db.WithContext(ctx).Find(&user).Error
	if err != nil {
		return nil, fmt.Errorf("userRepo.FindByID: %w", err)
	}
	return user, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, req *model.User) (*model.User, error) {
	// var user model.User
	if err := r.db.WithContext(ctx).Create(req).Error; err != nil {
		return nil, fmt.Errorf("userRepo.CreateUser: %w", err)
	}
	return req, nil
}
