package model

import (
	"time"
)

type TodoDTO struct {
	ID              uint64
	ActivityGroupID uint64
	Title           string
	IsActive        int
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
