package domain

import "github.com/victoramsantos/game-contest/usecases/usecasesdomain"

type CharacterUsecase interface {
	NewCharacter(name string, className string) (*Character, error)
	GetCharacterDetails(name string) (*usecasesdomain.CharacterDetails, error)
	ListCharacters() ([]usecasesdomain.CharacterStatus, error)
}

type GameUsecase interface {
	StartFight(nameA string, nameB string) ([]string, error)
}
