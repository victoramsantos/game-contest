package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victoramsantos/game-contest/controller/controllerdomain"
	"github.com/victoramsantos/game-contest/domain"
)

type gameController struct {
	usecase domain.GameUsecase
}

func InitGameController(e *echo.Echo, usecase domain.GameUsecase) {
	handler := &gameController{
		usecase: usecase,
	}

	e.POST("/game/start", handler.Start)
}

func (this *gameController) Start(ctx echo.Context) error {
	var request controllerdomain.GameRequest
	err := ctx.Bind(&request)

	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusUnprocessableEntity, controllerdomain.ResponseError{err.Error()})
	}

	logger, err := this.usecase.StartFight(request.CharacterA, request.CharacterB)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, controllerdomain.ResponseError{err.Error()})
	}

	return ctx.JSON(http.StatusOK, logger)
}
