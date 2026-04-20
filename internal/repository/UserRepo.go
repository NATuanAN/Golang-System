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
		return nil, err
	}
	return user, nil
}
