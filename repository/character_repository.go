package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/victoramsantos/game-contest/domain"
)

type characterRepository struct {
	classes    []domain.Class
	characters []domain.Character
}

func (this *characterRepository) GetClassByName(className string) (*domain.Class, error) {
	for _, class := range this.classes {
		if strings.EqualFold(class.Name, className) {
			return &class, nil
		}
	}
	return nil, errors.New("no class found with name")
}

func (this *characterRepository) GetCharacterByName(name string) (*domain.Character, error) {
	for _, character := range this.characters {
		if strings.EqualFold(character.Name, name) {
			return &character, nil
		}
	}
	return nil, errors.New("no character found with name")
}

func (this *characterRepository) CreateCharacter(character domain.Character) error {
	this.characters = append(this.characters, character)
	return nil
}

func (this *characterRepository) ListCharacters() ([]domain.Character, error) {
	return this.characters, nil
}

func NewCharacterRepository() domain.CharacterRepository {
	classes := make([]domain.Class, 0)
	//Warrior

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
	classes = append(classes, domain.Class{
		Name:       "Warrior",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	})

	//Thief
	attributes = &domain.Attributes{
		Life: domain.Attribute{
			Name:  "Life",
			Value: 15,
		},

		Strength: domain.Attribute{
			Name:  "Strength",
			Value: 4,
		},
		Skill: domain.Attribute{
			Name:  "Skill",
			Value: 10,
		},
		Intelligence: domain.Attribute{
			Name:  "Intelligence",
			Value: 4,
		},
	}

	attacks = make([]domain.AttributeImprovment, 0)
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Strength,
		Improvement:     25,
	})
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Skill,
		Improvement:     100,
	})
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Intelligence,
		Improvement:     25,
	})

	velocity = make([]domain.AttributeImprovment, 0)
	velocity = append(velocity, domain.AttributeImprovment{
		TargetAttribute: &attributes.Skill,
		Improvement:     80,
	})
	classes = append(classes, domain.Class{
		Name:       "Thief",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	})

	//Mage
	attributes = &domain.Attributes{
		Life: domain.Attribute{
			Name:  "Life",
			Value: 12,
		},

		Strength: domain.Attribute{
			Name:  "Strength",
			Value: 5,
		},
		Skill: domain.Attribute{
			Name:  "Skill",
			Value: 6,
		},
		Intelligence: domain.Attribute{
			Name:  "Intelligence",
			Value: 10,
		},
	}

	attacks = make([]domain.AttributeImprovment, 0)
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Strength,
		Improvement:     20,
	})
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Skill,
		Improvement:     50,
	})
	attacks = append(attacks, domain.AttributeImprovment{
		TargetAttribute: &attributes.Intelligence,
		Improvement:     150,
	})

	velocity = make([]domain.AttributeImprovment, 0)
	velocity = append(velocity, domain.AttributeImprovment{
		TargetAttribute: &attributes.Strength,
		Improvement:     20,
	})
	velocity = append(velocity, domain.AttributeImprovment{
		TargetAttribute: &attributes.Strength,
		Improvement:     20,
	})
	classes = append(classes, domain.Class{
		Name:       "Mage",
		Attributes: attributes,
		Attack:     attacks,
		Velocity:   velocity,
	})

	fmt.Println(classes)

	return &characterRepository{
		classes:    classes,
		characters: make([]domain.Character, 0),
	}
}
