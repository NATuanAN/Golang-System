package service_test

import (
	"context"
	"errors"
	"go-project/internal/model"
	"go-project/internal/service"
	mymock "go-project/mock"
	"go-project/pkg/apperror"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser_Success(t *testing.T) {
	t.Log("case: create user success")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	input := &model.User{
		Name:     "Nguyen Van A",
		Email:    "abc@gmail.com",
		Password: "plaintext123",
	}

	mockRepo.On("FindByEmail", context.Background(), input.Email).
		Return(nil, apperror.ErrNotFound)

	mockRepo.On("CreateUser", context.Background(),
		mock.MatchedBy(func(u *model.User) bool {
			return u.Email == input.Email
		}),
	).Return(&model.User{ID: 1, Email: input.Email}, nil)

	result, err := svc.CreateUser(context.Background(), input)
	t.Logf("result id: %d, err: %v", result.ID, err)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailConflict(t *testing.T) {
	t.Log("case: create user email conflict")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	input := &model.User{Email: "a@gmail.com", Password: "123"}

	mockRepo.On("FindByEmail", context.Background(), input.Email).
		Return(&model.User{Email: "a@gmail.com"}, nil)

	_, err := svc.CreateUser(context.Background(), input)
	t.Logf("err: %v", err)

	assert.ErrorIs(t, err, apperror.ErrConflict)
	mockRepo.AssertExpectations(t)
}

func TestGetById_Success(t *testing.T) {
	t.Log("case: get by id success")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	mockRepo.On("FindByID", context.Background(), "1").
		Return(&model.User{ID: 1, Email: "user@gmail.com"}, nil)

	result, err := svc.GetById(context.Background(), "1")
	t.Logf("result id: %d, err: %v", result.ID, err)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetById_NotFound(t *testing.T) {
	t.Log("case: get by id not found")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	mockRepo.On("FindByID", context.Background(), "404").
		Return(nil, apperror.ErrNotFound)

	result, err := svc.GetById(context.Background(), "404")
	t.Logf("result nil: %v, err: %v", result == nil, err)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, apperror.ErrNotFound)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	t.Log("case: login success")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("correctpass"), bcrypt.MinCost)
	mockRepo.On("FindByEmail", context.Background(), "nguyenvana4@gmail.com").
		Return(&model.User{ID: 1, Email: "nguyenvana4@gmail.com", Password: string(hashed)}, nil)

	token, err := svc.Login(context.Background(), "nguyenvana4@gmail.com", "correctpass")
	t.Logf("token empty: %v, err: %v", token == "", err)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_EmptyEmailOrPassword(t *testing.T) {
	t.Log("case: login empty email or password")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	_, err := svc.Login(context.Background(), "", "123")
	t.Logf("err (empty email): %v", err)
	assert.ErrorIs(t, err, apperror.ErrBadRequest)

	_, err = svc.Login(context.Background(), "nguyenvana4@gmail.com", "")
	t.Logf("err (empty password): %v", err)
	assert.ErrorIs(t, err, apperror.ErrBadRequest)

	mockRepo.AssertNotCalled(t, "FindByEmail")
}

func TestLogin_EmailNotFound(t *testing.T) {
	t.Log("case: login email not found")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	mockRepo.On("FindByEmail", context.Background(), "noone@gmail.com").
		Return(nil, apperror.ErrNotFound)

	_, err := svc.Login(context.Background(), "noone@gmail.com", "123")
	t.Logf("err: %v", err)

	assert.ErrorIs(t, err, apperror.ErrUnauthorized)
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	t.Log("case: login wrong password")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("correctpass"), bcrypt.MinCost)
	mockRepo.On("FindByEmail", context.Background(), "user@gmail.com").
		Return(&model.User{ID: 1, Email: "user@gmail.com", Password: string(hashed)}, nil)

	_, err := svc.Login(context.Background(), "user@gmail.com", "wrongpass")
	t.Logf("err: %v", err)

	assert.ErrorIs(t, err, apperror.ErrUnauthorized)
	mockRepo.AssertExpectations(t)
}

func TestLogin_DBError(t *testing.T) {
	t.Log("case: login db error")
	mockRepo := new(mymock.UserRepository)
	svc := service.NewUserService(mockRepo)

	mockRepo.On("FindByEmail", context.Background(), "user@gmail.com").
		Return(nil, errors.New("connection refused"))

	_, err := svc.Login(context.Background(), "user@gmail.com", "123")
	t.Logf("err: %v", err)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
