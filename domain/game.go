package domain

import (
	"math/rand"
	"time"
)

type GameLog interface {
	Start(*Game) string
	Attack(*Game) string
	Finish(*Game) string
}

type Game struct {
	Attacker *Player
	Opponent *Player
}

func NewGame(charA *Character, charB *Character) *Game {
	var first, second *Player
	var appliedVelocityA, appliedVelocityB int
	for {
		appliedVelocityA = gameVelocity(charA)
		appliedVelocityB = gameVelocity(charB)

		if appliedVelocityA == appliedVelocityB {
			continue
		}

		if appliedVelocityA > appliedVelocityB {
			first = NewPlayer(charA, appliedVelocityA, 0)
			second = NewPlayer(charB, appliedVelocityB, 0)
		} else if appliedVelocityB > appliedVelocityA {
			first = NewPlayer(charB, appliedVelocityB, 0)
			second = NewPlayer(charA, appliedVelocityA, 0)
		}
		break
	}

	return &Game{
		Attacker: first,
		Opponent: second,
	}
}

func (this *Game) Attack() {
	this.Opponent.Character.Class.Attributes.Life.Value -= this.Attacker.Attack
	if this.Opponent.Character.Class.Attributes.Life.Value <= 0 {
		this.Opponent.Character.Class.Attributes.Life.Value = 0
		this.Opponent.Character.IsAlive = false
	}
}

func (this *Game) SwitchPositions() {
	this.Attacker.UpdateAttack()
	this.Opponent.UpdateAttack()

	temp := this.Attacker
	this.Attacker = this.Opponent
	this.Opponent = temp
}

func (this *Game) DidGameFinished() bool {
	return this.Opponent.Character.Class.Attributes.Life.Value == 0
}

type Player struct {
	Character *Character
	Velocity  int
	Attack    int
}

func NewPlayer(character *Character, velocity int, attack int) *Player {
	return &Player{
		Character: character,
		Velocity:  velocity,
		Attack:    attack,
	}
}
func (this *Player) UpdateAttack() {
	randomizer := sanitizeRand()
	this.Attack = randomizer.Intn(this.Character.Class.GetAttack())
}

func gameVelocity(character *Character) int {
	randomizer := sanitizeRand()
	return randomizer.Intn(character.Class.GetVelocity())
}

func sanitizeRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
