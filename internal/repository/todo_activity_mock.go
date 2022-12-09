package repository

import (
	"ddd-to-do-list/internal/aggregate"

	"github.com/stretchr/testify/mock"
)

type TodoMock struct {
	mock.Mock
}

func (m *TodoMock) GetTodo() (todo aggregate.Todos, err error) {
	args := m.Called()

	return args.Get(0).(aggregate.Todos), args.Error(1)
}

func (m *TodoMock) CreateTodo(todo *aggregate.Todo) error {
	args := m.Called(todo)

	return args.Error(0)
}

func (m *TodoMock) GetTodoByID(id uint64) (res aggregate.Todos, err error) {
	args := m.Called(id)

	return args.Get(0).(aggregate.Todos), args.Error(1)
}
