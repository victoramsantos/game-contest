package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/victoramsantos/game-contest/domain"
	"github.com/victoramsantos/game-contest/repository"
	"github.com/victoramsantos/game-contest/repository/mocks"
	"github.com/victoramsantos/game-contest/usecases"
)

func TestStartFight(t *testing.T) {
	mockRepository := new(mocks.CharacterRepository)
	usecase := usecases.NewGameUsecase(mockRepository, repository.NewLogger())

	charNameA := "john"
	charNameB := "mary"

	attributes := &domain.Attributes{
		Life: domain.Attribute{
			Name:  "Life",
			Value: 20,
		},

		Strength: domain.Attribute{
			Name:  "Strength",
			Value: 10,
		},
		Skill: domain.Attribute{
			Name:  "Skill",
			Value: 5,
		},
		Intelligence: domain.Attribute{
			Name:  "Intelligence",
			Value: 5,
		},
	}

	attacks := make([]domain.AttributeImprovment, 0)
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Strength,
		Improvement:     80,
	})
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Skill,
		Improvement:     20,
	})

	velocity := make([]domain.AttributeImprovment, 0)
	velocity = append(velocity, domain.AttributeImprovment{
		TargetAttribute: &attributes.Skill,
		Improvement:     60,
	})
	velocity = append(velocity, domain.AttributeImprovment{
		TargetAttribute: &attributes.Intelligence,
		Improvement:     20,
	})

	warrior := domain.Class{
		Name:       "Warrior",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	}

	charA := &domain.Character{
		Name:    charNameA,
		Class:   &warrior,
		IsAlive: true,
	}

	charB := &domain.Character{
		Name:    charNameB,
		Class:   &warrior,
		IsAlive: true,
	}

	t.Run("success", func(t *testing.T) {
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(charA, nil)
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(charB, nil)
		mockRepository.On("UpdateCharacter", mock.AnythingOfType("*domain.Character")).Return(nil)
		mockRepository.On("UpdateCharacter", mock.AnythingOfType("*domain.Character")).Return(nil)

		logger, err := usecase.StartFight(charNameA, charNameB)

		assert.NoError(t, err)
		assert.NotNil(t, logger)
		assert.NotEmpty(t, logger)

		mockRepository.AssertExpectations(t)
	})

	t.Run("fight must be between different characters", func(t *testing.T) {
		logger, err := usecase.StartFight(charNameA, charNameA)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, logger)

		mockRepository.AssertExpectations(t)
	})

	t.Run("charA should exists in repository", func(t *testing.T) {
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(nil, errors.New("no character found with name"))

		logger, err := usecase.StartFight("NotFound", "someone")

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, logger)

		mockRepository.AssertExpectations(t)
	})

	t.Run("charB should exists in repository", func(t *testing.T) {
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(charA, nil)
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(nil, errors.New("no character found with name"))

		logger, err := usecase.StartFight("someone", "NotFound")

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, logger)

		mockRepository.AssertExpectations(t)
	})
}
