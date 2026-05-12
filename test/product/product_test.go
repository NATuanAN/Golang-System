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
)

func TestProductService_GetAll_Success(t *testing.T) {
	t.Log("case: product get all success")
	mockRepo := new(mymock.ProductRepository)
	svc := service.NewProductService(mockRepo)

	mockRepo.On("FindAll", context.Background()).Return([]model.Product{
		{ProductID: 1, ProductName: "Product A"},
		{ProductID: 2, ProductName: "Product B"},
	}, nil)

	result, err := svc.GetAll(context.Background())
	t.Logf("result len: %d, err: %v", len(result), err)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAll_Error(t *testing.T) {
	t.Log("case: product get all error")
	mockRepo := new(mymock.ProductRepository)
	svc := service.NewProductService(mockRepo)

	mockRepo.On("FindAll", context.Background()).Return(nil, errors.New("db error"))

	result, err := svc.GetAll(context.Background())
	t.Logf("result nil: %v, err: %v", result == nil, err)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetById_Success(t *testing.T) {
	t.Log("case: product get by id success")
	mockRepo := new(mymock.ProductRepository)
	svc := service.NewProductService(mockRepo)

	mockRepo.On("FindById", context.Background(), 1).
		Return(&model.Product{ProductID: 1, ProductName: "Product A"}, nil)

	result, err := svc.GetById(context.Background(), 1)
	t.Logf("result id: %d, err: %v", result.ProductID, err)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.ProductID)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetById_NotFound(t *testing.T) {
	t.Log("case: product get by id not found")
	mockRepo := new(mymock.ProductRepository)
	svc := service.NewProductService(mockRepo)

	mockRepo.On("FindById", context.Background(), 404).
		Return(nil, apperror.ErrNotFound)

	result, err := svc.GetById(context.Background(), 404)
	t.Logf("result nil: %v, err: %v", result == nil, err)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, apperror.ErrNotFound)
	mockRepo.AssertExpectations(t)
}
