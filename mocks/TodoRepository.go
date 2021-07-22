package mocks

import (
	context "context"

	domain "github.com/ancene/go-rest-echo/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

type TodoRepository struct {
	mock.Mock
}

func (_m *TodoRepository) FindByID(ctx context.Context, id string) (*domain.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Todo
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
