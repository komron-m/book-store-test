package internal

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id string) (*Store, error)

	Update(ctx context.Context, store *Store) error
}
