package entities

import "github.com/krios2146/simulation-go/internal"

type Creature interface {
	New(c internal.Coordinates) Entity
	InteractsWith() Entity
	GetCoordinates() internal.Coordinates
}

type Predator struct {
	Coordinates  internal.Coordinates
	Speed        uint8
	AttackRating uint8
}

type Herbivore struct {
	Coordinates internal.Coordinates
	Speed       uint8
	Health      uint8
}

func (p Predator) New(c internal.Coordinates) Entity {
	return &Predator{
		Coordinates:  c,
		Speed:        2,
		AttackRating: 50,
	}
}

func (p Predator) InteractsWith() Entity {
	return Herbivore{}
}

func (p Predator) GetCoordinates() internal.Coordinates {
	return p.Coordinates
}

func (p Predator) GetConsoleSprite() rune {
	// 🐺
	return '\U0001F43A'
}

func (h Herbivore) New(c internal.Coordinates) Entity {
	return &Herbivore{
		Coordinates: c,
		Speed:       1,
		Health:      100,
	}
}

func (h Herbivore) InteractsWith() Entity {
	return Grass{}
}

func (h Herbivore) GetCoordinates() internal.Coordinates {
	return h.Coordinates
}

func (h Herbivore) GetConsoleSprite() rune {
	// 🐑
	return '\U0001F411'
}
