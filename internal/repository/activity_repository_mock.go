package repository

import (
	"ddd-to-do-list/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type ActivityMock struct {
	mock.Mock
}

func (m *ActivityMock) GetActivity() (res aggregate.Activities, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Activities), args.Error(1)
}

func (m *ActivityMock) CreateActivity(email, title string) error {
	args := m.Called(email, title)

	return args.Error(0)
}

func (m *ActivityMock) GetActivityByID(id uint64) (res aggregate.Activities, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Activities), args.Error(1)
}

func (m *ActivityMock) UpdateActivity(id uint64, email, title string) error {
	args := m.Called(email, title)

	return args.Error(0)
}
