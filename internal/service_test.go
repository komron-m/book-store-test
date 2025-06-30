package internal

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSubtract_Behaviour_Succeeds(t *testing.T) {
	// init request
	ctx := context.TODO()
	storeID := uuid.NewString()
	req := SubtractBooksRequest{
		StoreID: storeID,
		Amount:  1_000,
	}

	// init mock
	mockRepo := newMockRepo()

	// init desired entity (aggregate)
	store := NewStore(storeID, 2_000)

	// set-up expectations
	mockRepo.On("Get", ctx, storeID).Return(store, nil)
	mockRepo.On("Update", ctx, store).Return(nil)

	// init service and call Subtract method
	srv := NewService(mockRepo)
	err := srv.SubtractBooks(ctx, req)

	// asserting use-case behaved as expected
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestSubtract_StateChange(t *testing.T) {
	// init request
	ctx := context.TODO()
	storeID := uuid.NewString()
	req := SubtractBooksRequest{
		StoreID: storeID,
		Amount:  1_000,
	}

	// init fake
	fakeRepo := newFakeRepo()

	// init desired entity (aggregate)
	store := NewStore(storeID, 2_000)
	fakeRepo.mapStore[storeID] = store

	// init service and call Subtract method
	srv := NewService(fakeRepo)
	err := srv.SubtractBooks(ctx, req)

	assert.NoError(t, err)

	// asserting state has changed
	expectedAvailableBooks := 2_000 - 1_000
	assert.Equal(t, expectedAvailableBooks, store.AvailableBooks)
}
