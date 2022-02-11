package usecases

import (
	"log"

	"github.com/victoramsantos/game-contest/domain"
)

type gameUsecase struct {
	repository domain.CharacterRepository
	log        domain.GameLog
}

func NewGameUsecase(repository domain.CharacterRepository, log domain.GameLog) domain.GameUsecase {
	return &gameUsecase{
		repository: repository,
		log:        log,
	}
}

func (this *gameUsecase) StartFight(nameA string, nameB string) ([]string, error) {
	charA, err := this.repository.GetCharacterByName(nameA)
	if err != nil {
		log.Println("No character found for", charA)
		return nil, err
	}
	charB, err := this.repository.GetCharacterByName(nameB)
	if err != nil {
		log.Println("No character found for", charB)
		return nil, err
	}
	game := domain.NewGame(charA, charB)
	logger := make([]string, 0)

	logger = append(logger, this.log.Start(game))

	for {
		game.Attack()
		logger = append(logger, this.log.Attack(game))
		if game.DidGameFinished() {
			break
		}
		game.SwitchPositions()
	}
	logger = append(logger, this.log.Finish(game))

	return logger, nil
}
