package internal

type Entity interface {
	New(Coordinates) Entity
	GetCoordinates() Coordinates
}

type Grass struct {
	Coordinates
}

type Rock struct {
	Coordinates
}

type Tree struct {
	Coordinates
}

func (g Grass) New(c Coordinates) Entity {
	return &Grass{Coordinates: c}
}

func (g Grass) GetCoordinates() Coordinates {
	return g.Coordinates
}

func (g Grass) GetConsoleSprite() rune {
	// 🍀
	return '\U0001F340'
}

func (r Rock) New(c Coordinates) Entity {
	return &Rock{Coordinates: c}
}

func (r Rock) GetCoordinates() Coordinates {
	return r.Coordinates
}

func (r Rock) GetConsoleSprite() rune {
	// 🪨
	return '\U0001FAA8'
}

func (t Tree) New(c Coordinates) Entity {
	return &Tree{Coordinates: c}
}

func (t Tree) GetCoordinates() Coordinates {
	return t.Coordinates
}

func (t Tree) GetConsoleSprite() rune {
	// 🌳
	return '\U0001F333'
}
