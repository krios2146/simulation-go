package entities

import (
	"fmt"

	"github.com/krios2146/simulation-go/internal/coordinates"
)

type Predator struct {
	Coordinates  coordinates.Coordinates
	Speed        uint8
	AttackRating uint8
}

type Herbivore struct {
	Coordinates coordinates.Coordinates
	Speed       uint8
	Health      uint8
}

type Creature interface {
	MakeMove()
}

func (p Predator) New(c coordinates.Coordinates) Entity {
	return &Predator{
		Coordinates:  c,
		Speed:        2,
		AttackRating: 50,
	}
}

func (p *Predator) MakeMove() {
	fmt.Println("Predator makes move")
}

func (p Predator) GetConsoleSprite() rune {
	// 🐺
	return '\U0001F43A'
}

func (h Herbivore) New(c coordinates.Coordinates) Entity {
	return &Herbivore{
		Coordinates: c,
		Speed:       1,
		Health:      100,
	}
}

func (h *Herbivore) MakeMove() {
	fmt.Println("Herbivore makes move")
}

func (h Herbivore) GetConsoleSprite() rune {
	// 🐑
	return '\U0001F411'
}
