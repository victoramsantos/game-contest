package usecases

import (
	"errors"
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
	game, err := this.setupGame(nameA, nameB)
	if err != nil {
		return nil, err
	}

	logger := make([]string, 0)

	logger = append(logger, this.log.Start(game))

	for {
		game.Attack()
		logger = append(logger, this.log.Attack(game))
		if game.DidGameFinished() {
			this.updateStats(game)
			break
		}
		game.SwitchPositions()
	}
	logger = append(logger, this.log.Finish(game))

	return logger, nil
}

func (this *gameUsecase) updateStats(game *domain.Game) {
	this.repository.UpdateCharacter(game.Attacker.Character)
	this.repository.UpdateCharacter(game.Opponent.Character)
}

func (this *gameUsecase) setupGame(nameA string, nameB string) (*domain.Game, error) {
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

	if !charA.IsAlive || !charB.IsAlive {
		log.Println("Some character is already dead")
		return nil, errors.New("some character is already dead")
	}

	return domain.NewGame(charA, charB), nil
}
