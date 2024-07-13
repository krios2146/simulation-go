package entities

import "github.com/krios2146/simulation-go/internal"

type Entity interface {
	New(internal internal.Coordinates) Entity
	GetCoordinates() internal.Coordinates
}

type Grass struct {
	internal.Coordinates
}

type Rock struct {
	internal.Coordinates
}

type Tree struct {
	internal.Coordinates
}

func (g Grass) New(c internal.Coordinates) Entity {
	return &Grass{Coordinates: c}
}

func (g Grass) GetCoordinates() internal.Coordinates {
	return g.Coordinates
}

func (g Grass) GetConsoleSprite() rune {
	// 🍀
	return '\U0001F340'
}

func (r Rock) New(c internal.Coordinates) Entity {
	return &Rock{Coordinates: c}
}

func (r Rock) GetCoordinates() internal.Coordinates {
	return r.Coordinates
}

func (r Rock) GetConsoleSprite() rune {
	// 🪨
	return '\U0001FAA8'
}

func (t Tree) New(c internal.Coordinates) Entity {
	return &Tree{Coordinates: c}
}

func (t Tree) GetCoordinates() internal.Coordinates {
	return t.Coordinates
}

func (t Tree) GetConsoleSprite() rune {
	// 🌳
	return '\U0001F333'
}
