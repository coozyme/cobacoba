package handler

import (
	"ddd-to-do-list/internal/shared"
	"ddd-to-do-list/internal/usecase"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase usecase.ActivityUsecase
}

func (h *handler) HandlerGetActivites(c echo.Context) error {
	activities, err := h.usecase.GetActivity()
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", nil, nil).JSON(c)

	}

	return shared.NewResponse(true, 200, "success", nil, ActivityResponse{}.Response(activities)).JSON(c)
}

func (h *handler) HandlerGetActivitesByUUID(c echo.Context) error {
	id := c.QueryParam("id")

	uintID, _ := strconv.ParseUint(id, 10, 64)
	activities, err := h.usecase.GetActivityByID(uintID)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)

	}

	return shared.NewResponse(true, 200, "success", ActivityResponse{}.Response(activities), nil).JSON(c)
}

func (h *handler) HandlerCreateActivity(c echo.Context) error {
	var body ReqCreateActivity
	c.Bind(&body)
	// json.NewDecoder(req.Body).Decode(&body)
	log.Println("LOGC-RE", body.Email)
	err := h.usecase.CreateActivity(body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}

	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func (h *handler) HandlerUpdateActivity(c echo.Context) error {
	var body ReqUpdateActivity
	c.Bind(&body)

	err := h.usecase.UpdateActivity(body.ID, body.Email, body.Title)
	if err != nil {
		log.Println(err)

		return shared.NewResponse(false, 400, "failed", err.Error(), nil).JSON(c)
	}

	return shared.NewResponse(true, 200, "success", nil, nil).JSON(c)
}

func NewHandler(usecase usecase.ActivityUsecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
