package internal

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	repo           Repository
	holidayChecker HolidayChecker
}

func NewService(repo Repository, checker HolidayChecker) *Service {
	return &Service{
		repo:           repo,
		holidayChecker: checker,
	}
}

type SubtractBooksRequest struct {
	StoreID string `json:"store_id"`
	Amount  int    `json:"amount"`
}

func (s *Service) SubtractBooks(ctx context.Context, req SubtractBooksRequest) error {
	store, err := s.repo.Get(ctx, req.StoreID)
	if err != nil {
		return fmt.Errorf("get book store: %w", err)
	}

	today := time.Now()
	if isHoliday := s.holidayChecker.Check(today); isHoliday {
		return fmt.Errorf("do not subtract on holidays")
	}

	if err := store.Subtract(req.Amount); err != nil {
		return fmt.Errorf("subtract from store: %w", err)
	}

	if err := s.repo.Update(ctx, store); err != nil {
		return fmt.Errorf("update store: %w", err)
	}

	return nil
}
