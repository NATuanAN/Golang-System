package repository

import (
	"context"
	"errors"
	"fmt"
	"go-project/internal/model"
	"go-project/pkg/apperror"

	"golang.org/x/crypto/bcrypt"
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

func (r *userRepo) CreateUser(ctx context.Context, req *model.User) (*model.User, error) {
	repo := r.db.WithContext(ctx)

	var existed model.User
	err := repo.First(&existed, "email = ?", req.Email).Error

	if err == nil {
		return nil, apperror.ErrConflict
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("userRepo.Register check email: %w", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("userRepo.Register hash password: %w", err)
	}
	req.Password = string(hashedPassword)
	if err := repo.Create(&req).Error; err != nil {
		return nil, fmt.Errorf("userRepo.Register create: %w", err)
	}

	return req, nil
}
