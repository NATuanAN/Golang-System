// internal/service/user_service_test.go
package service_test

import (
	"context"
	"go-project/internal/model"
	"go-project/internal/service"
	"go-project/mock"
	"go-project/pkg/apperror"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mock.UserRepository)
	svc := service.NewUserService(mockRepo)

	input := &model.User{
		Name:     "Nguyen Van A",
		Email:    "a@gmail.com",
		Password: "plaintext123",
	}

	// FindByEmail trả nil, nil — email chưa tồn tại
	mockRepo.On("FindByEmail", context.Background(), input.Email).
		Return(nil, nil)

	// CreateUser trả về user đã có ID
	mockRepo.On("CreateUser", context.Background(), input).
		Return(&model.User{ID: 1, Email: input.Email}, nil)

	result, err := svc.CreateUser(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailConflict(t *testing.T) {
	mockRepo := new(mock.UserRepository)
	svc := service.NewUserService(mockRepo)

	input := &model.User{Email: "a@gmail.com", Password: "123"}

	// FindByEmail trả về user — email đã tồn tại
	mockRepo.On("FindByEmail", context.Background(), input.Email).
		Return(&model.User{Email: "a@gmail.com"}, nil)

	_, err := svc.CreateUser(context.Background(), input)

	assert.ErrorIs(t, err, apperror.ErrConflict)
	mockRepo.AssertExpectations(t)
}
