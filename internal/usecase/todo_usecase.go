package usecase

import "ddd-to-do-list/internal/aggregate"

type TodoUsecase interface {
	GetTodo() (aggregate.Todos, error)
	GetTodoByID(id uint64) (aggregate.Todos, error)
	CreateTodo(activitGroupID uint64, titile string) error
	UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error
}
