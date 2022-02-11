package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/victoramsantos/game-contest/domain"
)

type CharacterRepository struct {
	mock.Mock
}

func (_mock *CharacterRepository) GetClassByName(className string) (*domain.Class, error) {
	ret := _mock.Called()
	return ret.Get(0).(*domain.Class), nil
}

func (_mock *CharacterRepository) GetCharacterByName(name string) (*domain.Character, error) {
	ret := _mock.Called()
	return ret.Get(0).(*domain.Character), nil
}

func (_mock *CharacterRepository) CreateCharacter(character domain.Character) error {
	_mock.Called(character)
	return nil
}

func (_mock *CharacterRepository) ListCharacters() ([]domain.Character, error) {
	ret := _mock.Called()
	return ret.Get(0).([]domain.Character), nil
}
