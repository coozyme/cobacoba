package router

import (
	"ddd-to-do-list/internal/delivery/handler"
	"ddd-to-do-list/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Router(route *echo.Echo, usecase usecase.ActivityUsecase) {
	h := handler.NewHandler(usecase)

	v1 := route.Group("v1")
	{
		v1.GET("/list-todo", h.HandlerGetActivites)
		v1.GET("", h.HandlerGetActivites)
		// v1.GET("/id", h.HandlerGetActivitesByUUID)
		v1.POST("/createActivity", h.HandlerCreateActivity)
		v1.PUT("/updateActivity", h.HandlerUpdateActivity)
	}

}
