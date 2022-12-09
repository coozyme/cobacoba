package usecase

import "ddd-to-do-list/internal/aggregate"

type ActivityUsecase interface {
	GetActivity() (aggregate.Activities, error)
	GetActivityByID(id uint64) (aggregate.Activities, error)
	CreateActivity(email, titile string) error
	UpdateActivity(id uint64, email, title string) error
}
