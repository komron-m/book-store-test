package internal

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	mockedHolidayChecker := newMockHolidayChecker()

	// init desired entity (aggregate)
	store := NewStore(storeID, 2_000)

	// set-up expectations
	mockRepo.On("Get", ctx, storeID).Return(store, nil)
	mockRepo.On("Update", ctx, store).Return(nil)
	mockedHolidayChecker.On("Check", mock.AnythingOfType("time.Time")).Return(false)

	// init service and call Subtract method
	srv := NewService(mockRepo, mockedHolidayChecker)
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
	srv := NewService(fakeRepo, dummyHolidayChecker{})
	err := srv.SubtractBooks(ctx, req)

	assert.NoError(t, err)

	// asserting state has changed
	expectedAvailableBooks := 2_000 - 1_000
	assert.Equal(t, expectedAvailableBooks, store.AvailableBooks)
}
