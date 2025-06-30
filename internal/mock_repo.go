package internal

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func newMockRepo() *mockRepository {
	return new(mockRepository)
}

func (r *mockRepository) Get(ctx context.Context, id string) (*Store, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*Store), args.Error(1)
}

func (r *mockRepository) Update(ctx context.Context, store *Store) error {
	args := r.Called(ctx, store)
	return args.Error(0)
}
