package handler

import (
	"ddd-to-do-list/internal/aggregate"
)

type ActivityResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Title string `json:"title"`
}

func (response ActivityResponse) Response(activity aggregate.Activities) ActivityResponse {
	for _, src := range activity {
		response.ID = src.ID
		response.Email = src.Email
		response.Title = src.Title
	}

	return ActivityResponse{
		response.ID,
		response.Email,
		response.Title,
	}
}
