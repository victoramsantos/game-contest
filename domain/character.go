package domain

import (
	"fmt"

	"github.com/victoramsantos/game-contest/usecases/usecasesdomain"
)

type Character struct {
	Name    string `json:"name" vd:"len($)>1 && len($) <= 15 && regexp('([A-Za-z_])+')"`
	Class   *Class `json:"class"`
	IsAlive bool   `json:"is_alive"`
}

func (this Character) ToDetails() *usecasesdomain.CharacterDetails {
	var attack, velocity string

	for _, mod := range this.Class.Attack {
		attack += fmt.Sprintf("%v%% de %v", mod.Improvement, mod.TargetAttribute.Name) + "+"
	}
	attack = attack[:len(attack)-1]

	for _, mod := range this.Class.Velocity {
		velocity += fmt.Sprintf("%v%% de %v", mod.Improvement, mod.TargetAttribute.Name) + "+"
	}
	velocity = velocity[:len(velocity)-1]

	return &usecasesdomain.CharacterDetails{
		Name:         this.Name,
		Class:        this.Class.Name,
		Life:         this.Class.Attributes.Life.Name,
		Strength:     this.Class.Attributes.Strength.Name,
		Skill:        this.Class.Attributes.Skill.Name,
		Intelligence: this.Class.Attributes.Intelligence.Name,
		Attack:       attack,
		Velocity:     velocity,
	}
}

func (this Character) ToStatus() *usecasesdomain.CharacterStatus {
	return &usecasesdomain.CharacterStatus{
		Name:    this.Name,
		Class:   this.Class.Name,
		IsAlive: this.IsAlive,
	}
}

type Class struct {
	Name       string                `json:"name"`
	Attributes *Attributes           `json:"attributes"`
	Attack     []AttributeImprovment `json:"attack_improvement"`
	Velocity   []AttributeImprovment `json:"velocity_improvement"`
}

type Attributes struct {
	Life         Attribute `json:"life"`
	Strength     Attribute `json:"strength"`
	Skill        Attribute `json:"skill"`
	Intelligence Attribute `json:"intelligence"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value uint8  `json:"value"`
}

type AttributeImprovment struct {
	TargetAttribute *Attribute `json:"target_attribute"`
	Improvement     uint16     `json:"improvement"`
}

type CharacterUsecase interface {
	NewCharacter(name string, className string) (*Character, error)
	GetCharacterDetails(name string) (*usecasesdomain.CharacterDetails, error)
	ListCharacters() ([]usecasesdomain.CharacterStatus, error)
}

type CharacterRepository interface {
	GetClassByName(className string) (*Class, error)
	GetCharacterByName(name string) (*Character, error)
	CreateCharacter(character Character) error
	ListCharacters() ([]Character, error)
}
