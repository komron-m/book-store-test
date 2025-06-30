package internal

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type HolidayChecker interface {
	Check(dt time.Time) bool // equivalent to IsHoliday(dt time.Time) (yes/no flag)
}

type mockHolidayChecker struct {
	mock.Mock
}

func newMockHolidayChecker() *mockHolidayChecker {
	return new(mockHolidayChecker)
}

func (h *mockHolidayChecker) Check(dt time.Time) bool {
	args := h.Called(dt)
	return args.Bool(0)
}

type dummyHolidayChecker struct{}

func (dummyHolidayChecker) Check(_ time.Time) bool { return false }
