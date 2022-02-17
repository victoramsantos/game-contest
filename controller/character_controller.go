package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victoramsantos/game-contest/controller/controllerdomain"
	"github.com/victoramsantos/game-contest/domain"
)

type characterController struct {
	usecase domain.CharacterUsecase
}

func InitCharacterController(e *echo.Echo, usecase domain.CharacterUsecase) {
	handler := &characterController{
		usecase: usecase,
	}

	e.POST("/character", handler.CreateCharacter)
	e.GET("/character", handler.ListCharacters)
	e.GET("/character/:name", handler.GetCharacterDetails)
}

func (this *characterController) CreateCharacter(ctx echo.Context) error {
	controllerdomain.Simulation("character")
	var request controllerdomain.CharacterRequest
	err := ctx.Bind(&request)
	//TODO: Validate input

	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusUnprocessableEntity, controllerdomain.ResponseError{err.Error()})
	}

	character, err := this.usecase.NewCharacter(request.CharacterName, request.ClassName)
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, controllerdomain.ResponseError{err.Error()})
	}

	return ctx.JSON(http.StatusCreated, character)
}

func (this *characterController) GetCharacterDetails(ctx echo.Context) error {
	controllerdomain.Simulation("character")
	character, err := this.usecase.GetCharacterDetails(ctx.Param("name"))
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, controllerdomain.ResponseError{err.Error()})
	}

	return ctx.JSON(http.StatusOK, character)
}

func (this *characterController) ListCharacters(ctx echo.Context) error {
	controllerdomain.Simulation("character")
	characters, err := this.usecase.ListCharacters()
	if err != nil {
		log.Println(err.Error())
		return ctx.JSON(http.StatusInternalServerError, controllerdomain.ResponseError{err.Error()})
	}

	return ctx.JSON(http.StatusOK, characters)
}
