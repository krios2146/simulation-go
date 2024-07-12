package entities

import (
	"fmt"

	"github.com/krios2146/simulation-go/internal/coordinates"
)

type Predator struct {
	Coordinates  coordinates.Coordinates
	Speed        uint8
	AttackRating uint8
	Sprite       rune
}

type Herbivore struct {
	Coordinates coordinates.Coordinates
	Speed       uint8
	Health      uint8
	Sprite      rune
}

type Creature interface {
	MakeMove()
}

func (p Predator) CanMove() bool {
	return true
}

func (p *Predator) MakeMove() {
	fmt.Println("Predator makes move")
}

func (p Predator) GetConsoleSprite() rune {
	// 🐺
	return '\U0001F43A'
}

func (p Predator) String() string {
	return fmt.Sprintf("predator: { coordinates: %v, speed: %v, attackRate: %v }",
		p.Coordinates, p.Speed, p.AttackRating,
	)
}

func (h Herbivore) CanMove() bool {
	return true
}

func (h *Herbivore) MakeMove() {
	fmt.Println("Herbivore makes move")
}

func (h Herbivore) GetConsoleSprite() rune {
	// 🐑
	return '\U0001F411'
}

func (h Herbivore) String() string {
	return fmt.Sprintf("coordinates: { coordinates: %v, speed: %v, health: %v }",
		h.Coordinates, h.Speed, h.Health,
	)
}
