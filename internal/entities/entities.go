package entities

import (
	"github.com/krios2146/simulation-go/internal/coordinates"
)

type Entity interface {
	CanMove() bool // Effectively indicates whether Entity is Creature or not
}

type Grass struct {
	coordinates.Coordinates
	health uint8
}

type Rock struct {
	coordinates.Coordinates
}

type Tree struct {
	coordinates.Coordinates
}

func (g Grass) CanMove() bool {
	return false
}

func (g Grass) GetConsoleSprite() rune {
	// 🍀
	return '\U0001F340'
}

func (r Rock) CanMove() bool {
	return false
}

func (r Rock) GetConsoleSprite() rune {
	// 🪨
	return '\U0001FAA8'
}

func (t Tree) CanMove() bool {
	return false
}

func (t Tree) GetConsoleSprite() rune {
	// 🌳
	return '\U0001F333'
}
