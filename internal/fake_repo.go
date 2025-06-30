package internal

import (
	"context"
	"fmt"
)

type fakeRepository struct {
	mapStore map[string]*Store
}

func newFakeRepo() *fakeRepository {
	ms := make(map[string]*Store)
	return &fakeRepository{ms}
}

func (r *fakeRepository) Get(_ context.Context, id string) (*Store, error) {
	store, ok := r.mapStore[id]
	if ok {
		return store, nil
	}
	return nil, fmt.Errorf("not found")
}

func (r *fakeRepository) Update(_ context.Context, store *Store) error {
	r.mapStore[store.ID] = store
	return nil
}
