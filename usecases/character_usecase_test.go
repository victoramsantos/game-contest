package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/victoramsantos/game-contest/domain"
	"github.com/victoramsantos/game-contest/repository/mocks"
	"github.com/victoramsantos/game-contest/usecases"
)

func TestNewCharacter(t *testing.T) {
	mockRepository := new(mocks.CharacterRepository)
	usecase := usecases.NewCharacterUsecase(mockRepository)

	characterName := "john"
	className := "Warrior"

	expectedClass1 := &domain.Class{
		Name:       className,
		Attributes: nil,
		Attack:     nil,
		Velocity:   nil,
	}

	expectedCharacter1 := domain.Character{
		Name:    characterName,
		Class:   expectedClass1,
		IsAlive: true,
	}

	t.Run("success", func(t *testing.T) {
		mockRepository.On("CreateCharacter", mock.AnythingOfType("domain.Character")).Return(expectedCharacter1).Once()
		mockRepository.On("GetClassByName", mock.AnythingOfType("string")).Return(expectedClass1, nil).Once()

		character, err := usecase.NewCharacter("john", "Warrior")

		assert.NoError(t, err)
		assert.Equal(t, expectedCharacter1.Name, character.Name)
		assert.Equal(t, expectedCharacter1.Class.Name, character.Class.Name)
		assert.Equal(t, characterName, character.Name)
		assert.Equal(t, className, character.Class.Name)
		assert.True(t, character.IsAlive)

		mockRepository.AssertExpectations(t)
	})
}
