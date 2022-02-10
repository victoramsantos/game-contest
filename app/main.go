package main

import (
	"log"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	controller "github.com/victoramsantos/game-contest/controller"
	repository "github.com/victoramsantos/game-contest/repository"
	usecases "github.com/victoramsantos/game-contest/usecases"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	initPrometheus(e)
	initCharacterService(e)

	log.Fatal(e.Start(":" + viper.GetString("app.server.port")))
}

func initPrometheus(e *echo.Echo) {
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}

func initCharacterService(e *echo.Echo) {
	repository := repository.NewCharacterRepository()
	usecase := usecases.NewCharacterUsecase(repository)
	controller.InitCharacterController(e, usecase)
}

func init() {
	setupEnvVars()
}

func setupEnvVars() {
	environment, isSet := os.LookupEnv("ENVIRONMENT")
	if !isSet {
		environment = "local"
	}
	viper.SetConfigName(environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
