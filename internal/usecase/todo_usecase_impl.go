package usecase

import (
	"ddd-to-do-list/internal/aggregate"
	"ddd-to-do-list/internal/repository"
)

type todoUsecase struct {
	repo repository.TodoRepository
}

func (u *todoUsecase) GetTodo() (aggregate.Todos, error) {
	todo, err := u.repo.GetTodo()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) GetTodoByID(id uint64) (aggregate.Todos, error) {
	activity, err := u.repo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	return activity, nil
}

func (u *todoUsecase) CreateTodo(activitGroupID uint64, titile string) error {
	err := u.repo.CreateTodo(activitGroupID, titile)
	if err != nil {
		return err
	}
	return nil
}

func (u *todoUsecase) UpdateTodo(id uint64, activitGroupID, IsActive int, title, priority string) error {
	err := u.repo.UpdateTodo(id, activitGroupID, IsActive, title, priority)
	if err != nil {
		return err
	}

	return nil
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		repo: repo,
	}
}
