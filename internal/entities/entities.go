package entities

import (
	"github.com/krios2146/simulation-go/internal/coordinates"
)

type Entity interface {
	New(coordinates coordinates.Coordinates) Entity
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

func (g Grass) New(c coordinates.Coordinates) Entity {
	return &Grass{Coordinates: c}
}

func (g Grass) GetConsoleSprite() rune {
	// 🍀
	return '\U0001F340'
}

func (r Rock) New(c coordinates.Coordinates) Entity {
	return &Rock{Coordinates: c}
}

func (r Rock) GetConsoleSprite() rune {
	// 🪨
	return '\U0001FAA8'
}

func (t Tree) New(c coordinates.Coordinates) Entity {
	return &Tree{Coordinates: c}
}

func (t Tree) GetConsoleSprite() rune {
	// 🌳
	return '\U0001F333'
}
