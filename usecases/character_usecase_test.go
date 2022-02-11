package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/victoramsantos/game-contest/domain"
	"github.com/victoramsantos/game-contest/repository/mocks"
	"github.com/victoramsantos/game-contest/usecases"
	"github.com/victoramsantos/game-contest/usecases/usecasesdomain"
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

	// TODO
	// t.Run("character already exists", func(t *testing.T) {
	// 	mockRepository.On("CreateCharacter", mock.AnythingOfType("domain.Character")).Return(expectedCharacter1).Once()
	// 	mockRepository.On("GetClassByName", mock.AnythingOfType("string")).Return(expectedClass1, nil).Once()

	// 	usecase.NewCharacter("john", "Warrior")
	// 	character, err := usecase.NewCharacter("john", "Warrior")

	// 	assert.Error(t, err)
	// 	assert.Nil(t, character)

	// 	mockRepository.AssertExpectations(t)
	// })

	t.Run("invalid name", func(t *testing.T) {
		mockRepository.On("CreateCharacter", mock.AnythingOfType("domain.Character")).Return(expectedCharacter1)
		mockRepository.On("GetClassByName", mock.AnythingOfType("string")).Return(expectedClass1, nil)

		character, err := usecase.NewCharacter("john1", "Warrior")
		assert.Error(t, err)
		assert.Nil(t, character)

		character, err = usecase.NewCharacter("1john", "Warrior")
		assert.Error(t, err)
		assert.Nil(t, character)

		character, err = usecase.NewCharacter("jo1hn", "Warrior")
		assert.Error(t, err)
		assert.Nil(t, character)

		character, err = usecase.NewCharacter("", "Warrior")
		assert.Error(t, err)
		assert.Nil(t, character)

		character, err = usecase.NewCharacter("aaaaaaaaaaaaaaaa", "Warrior")
		assert.Error(t, err)
		assert.Nil(t, character)

		mockRepository.AssertExpectations(t)
	})

	t.Run("valid name", func(t *testing.T) {
		mockRepository.On("CreateCharacter", mock.AnythingOfType("domain.Character")).Return(expectedCharacter1)
		mockRepository.On("GetClassByName", mock.AnythingOfType("string")).Return(expectedClass1, nil)

		character, err := usecase.NewCharacter("john_", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		character, err = usecase.NewCharacter("_john", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		character, err = usecase.NewCharacter("jo_hn", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		character, err = usecase.NewCharacter("a", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		character, err = usecase.NewCharacter("_", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		character, err = usecase.NewCharacter("aaaaaaaaaaaaaaa", "Warrior")
		assert.NoError(t, err)
		assert.NotNil(t, character)

		mockRepository.AssertExpectations(t)
	})

	t.Run("invalid class name", func(t *testing.T) {
		mockRepository.On("CreateCharacter", mock.AnythingOfType("domain.Character")).Return(expectedCharacter1)
		mockRepository.On("GetClassByName", mock.AnythingOfType("string")).Return(nil, errors.New("class not found"))

		character, err := usecase.NewCharacter("john", "Undefined")
		assert.Error(t, err)
		assert.Nil(t, character)

		mockRepository.AssertExpectations(t)
	})

}

func TestGetCharacterDetails(t *testing.T) {
	mockRepository := new(mocks.CharacterRepository)
	usecase := usecases.NewCharacterUsecase(mockRepository)

	characterName := "john"
	className := "Warrior"

	expectedDetailed1 := usecasesdomain.CharacterDetails{
		Name:         "john",
		Class:        "Warrior",
		Life:         "",
		Strength:     "",
		Skill:        "",
		Intelligence: "",
		Attack:       "",
		Velocity:     "",
	}

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

	class := domain.Class{
		Name:       "Warrior",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	}

	expectedCharacter1 := &domain.Character{
		Name:    "john",
		Class:   &class,
		IsAlive: true,
	}

	t.Run("success", func(t *testing.T) {
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(expectedCharacter1).Once()

		characterDetailed, err := usecase.GetCharacterDetails("john")

		assert.NoError(t, err)
		assert.Equal(t, expectedDetailed1.Name, characterDetailed.Name)
		assert.Equal(t, expectedDetailed1.Class, characterDetailed.Class)
		assert.Equal(t, characterName, characterDetailed.Name)
		assert.Equal(t, className, characterDetailed.Class)

		mockRepository.AssertExpectations(t)
	})

	t.Run("character not found", func(t *testing.T) {
		mockRepository.On("GetCharacterByName", mock.AnythingOfType("string")).Return(expectedCharacter1).Once()

		characterDetailed, err := usecase.GetCharacterDetails("NotFound")

		assert.Error(t, err)
		assert.Nil(t, characterDetailed)

		mockRepository.AssertExpectations(t)
	})
}

func TestListCharacters(t *testing.T) {
	mockRepository := new(mocks.CharacterRepository)
	usecase := usecases.NewCharacterUsecase(mockRepository)

	characterDetailed := usecasesdomain.CharacterStatus{
		Name:    "john",
		Class:   "Warrior",
		IsAlive: true,
	}
	expectedDetailedList := []usecasesdomain.CharacterStatus{characterDetailed}

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

	class := domain.Class{
		Name:       "Warrior",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	}

	character := &domain.Character{
		Name:    "john",
		Class:   &class,
		IsAlive: true,
	}

	list1 := []domain.Character{*character}

	t.Run("success", func(t *testing.T) {
		mockRepository.On("ListCharacters", mock.AnythingOfType("string")).Return(list1).Once()

		characters, err := usecase.ListCharacters()

		assert.NoError(t, err)
		assert.NotEmpty(t, characters)
		assert.Equal(t, expectedDetailedList[0].Name, characters[0].Name)
		assert.Equal(t, expectedDetailedList[0].Class, characters[0].Class)
		assert.Equal(t, expectedDetailedList[0].IsAlive, characters[0].IsAlive)

		mockRepository.AssertExpectations(t)
	})
}
