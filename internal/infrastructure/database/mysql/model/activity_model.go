package model

import (
	"ddd-to-do-list/internal/aggregate"
	"time"
)

type ActivityDTO struct {
	ID        uint64
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a ActivityDTO) Aggregate() *aggregate.Activity {
	activity, _ := aggregate.NewActivity(a.Email, a.Title)
	return activity
}
