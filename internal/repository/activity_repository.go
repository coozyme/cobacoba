package repository

import "ddd-to-do-list/internal/aggregate"

type ActivityRepository interface {
	GetActivity() (res aggregate.Activities, err error)
	CreateActivity(email, titile string) error
	GetActivityByID(id uint64) (res aggregate.Activities, err error)
	UpdateActivity(id uint64, email, title string) error
}
