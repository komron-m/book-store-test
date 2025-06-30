package internal

import "fmt"

type Store struct {
	ID string

	AvailableBooks int
}

func NewStore(id string, availableBooks int) *Store {
	return &Store{
		ID:             id,
		AvailableBooks: availableBooks,
	}
}

func (s *Store) Subtract(amount int) error {
	if s.AvailableBooks-amount < 0 {
		return fmt.Errorf("not enought books in store")
	}

	s.AvailableBooks -= amount
	return nil
}
