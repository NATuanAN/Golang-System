package service_test

import (
	"context"
	"errors"
	"go-project/internal/model"
	"go-project/internal/service"
	mymock "go-project/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderService_GetAll_Success(t *testing.T) {
	t.Log("case: order get all success")
	mockRepo := new(mymock.OrderRepository)
	svc := service.NewOrderService(mockRepo)

	mockRepo.On("FindAll", context.Background()).Return([]model.Order{
		{OrderID: "order-1"},
		{OrderID: "order-2"},
	}, nil)

	result, err := svc.GetAll(context.Background())
	t.Logf("result len: %d, err: %v", len(result), err)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestOrderService_GetAll_Error(t *testing.T) {
	t.Log("case: order get all error")
	mockRepo := new(mymock.OrderRepository)
	svc := service.NewOrderService(mockRepo)

	mockRepo.On("FindAll", context.Background()).Return(nil, errors.New("db error"))

	result, err := svc.GetAll(context.Background())
	t.Logf("result nil: %v, err: %v", result == nil, err)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
