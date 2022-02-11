package domain

import (
	"fmt"

	"github.com/victoramsantos/game-contest/usecases/usecasesdomain"
)

type Character struct {
	Name    string `json:"name" vd:"len($)>0 && len($) <= 15 && regexp('^([A-Za-z_])+$')"`
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

func (this Class) GetVelocity() int {
	velocity := 0
	for _, attr := range this.Velocity {
		velocity += attr.GetImprovement()
	}
	return velocity
}

func (this Class) GetAttack() int {
	attack := 0
	for _, attr := range this.Attack {
		attack += attr.GetImprovement()
	}
	return attack
}

type Attributes struct {
	Life         Attribute `json:"life"`
	Strength     Attribute `json:"strength"`
	Skill        Attribute `json:"skill"`
	Intelligence Attribute `json:"intelligence"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type AttributeImprovment struct {
	TargetAttribute *Attribute `json:"target_attribute"`
	Improvement     uint16     `json:"improvement"`
}

func (this AttributeImprovment) GetImprovement() int {
	return int(this.Improvement) * int(this.TargetAttribute.Value) / 100
}
