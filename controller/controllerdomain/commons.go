package controllerdomain

import (
	"log"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

type ResponseError struct {
	Message string `json:"message"`
}

func randomizer() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Simulation(api string) {
	if viper.GetBool("app.simulation.enabled") {
		maxSleepTime := 0
		if api == "game" {
			maxSleepTime = viper.GetInt("app.simulation.apis_sleep.game")
		} else if api == "character" {
			maxSleepTime = viper.GetInt("app.simulation.apis_sleep.character")
		}

		sleepTime := time.Duration(randomizer().Intn(maxSleepTime)) * time.Millisecond
		time.Sleep(sleepTime)
		log.Println("sleeping ", sleepTime, " Millisecond")
	}
}
