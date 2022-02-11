package repository

import (
	"fmt"

	"github.com/victoramsantos/game-contest/domain"
)

type gameLog struct {
}

func NewLogger() domain.GameLog {
	return &gameLog{}
}

func (this *gameLog) Start(game *domain.Game) string {
	return fmt.Sprintf("%v (%d) foi mais veloz que o %v (%d), e irá começar!", game.Attacker.Character.Name, game.Attacker.Velocity, game.Opponent.Character.Name, game.Opponent.Velocity)
}

func (this *gameLog) Attack(game *domain.Game) string {
	return fmt.Sprintf("%v atacou %v com %d de dano, %v com %d pontos de vida restantes", game.Attacker.Character.Name, game.Opponent.Character.Name, game.Attacker.Attack, game.Opponent.Character.Name, game.Opponent.Character.Class.Attributes.Life.Value)
}

func (this *gameLog) Finish(game *domain.Game) string {
	return fmt.Sprintf("%v venceu a batalha! %v ainda tem %d pontos de vida restantes!", game.Attacker.Character.Name, game.Attacker.Character.Name, game.Attacker.Character.Class.Attributes.Life.Value)
}
