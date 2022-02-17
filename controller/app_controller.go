package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victoramsantos/game-contest/controller/controllerdomain"
)

type appController struct {
}

func InitAppController(e *echo.Echo) {
	handler := &appController{}

	e.GET("/health", handler.Health)
}

func (this *appController) Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, controllerdomain.Health{
		Status: true,
	})
}
