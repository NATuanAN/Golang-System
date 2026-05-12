package mock

import (
	"context"
	"go-project/internal/model"

	"github.com/stretchr/testify/mock"
)

type OrderRepository struct {
	mock.Mock
}

func (m *OrderRepository) FindAll(ctx context.Context) ([]model.Order, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Order), args.Error(1)
}
