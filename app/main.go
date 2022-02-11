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

	characterRepository := repository.NewCharacterRepository()
	characterUsecase := usecases.NewCharacterUsecase(characterRepository)
	controller.InitCharacterController(e, characterUsecase)

	loggerRepository := repository.NewLogger()
	gameUsecase := usecases.NewGameUsecase(characterRepository, loggerRepository)
	controller.InitGameController(e, gameUsecase)

	log.Fatal(e.Start(":" + viper.GetString("app.server.port")))
}

func initPrometheus(e *echo.Echo) {
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
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
