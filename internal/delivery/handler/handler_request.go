package handler

type ReqCreateActivity struct {
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ReqUpdateActivity struct {
	ID    uint64 `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ReqCreateTodo struct {
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type ReqUpdateTodo struct {
	ID    uint64 `json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
	Title string `json:"title" validate:"required"`
}
