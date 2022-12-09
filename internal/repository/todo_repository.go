package repository

import "ddd-to-do-list/internal/aggregate"

type TodoRepository interface {
	GetTodo() (todo aggregate.Todos, err error)
	CreateTodo(activitGroupID uint64, title string) error
	GetTodoByID(id uint64) (res aggregate.Todos, err error)
	UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error
	DeleteTodo(id uint64) error
}
