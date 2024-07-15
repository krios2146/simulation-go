package internal

type Entity interface {
	GetCoordinates() Coordinates
	Spawnable
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
	return Grass{Coordinates: c}
}

func (g Grass) GetCoordinates() Coordinates {
	return g.Coordinates
}

func (g Grass) GetConsoleSprite() string {
	return "🍀"
}

func (r Rock) New(c Coordinates) Entity {
	return Rock{Coordinates: c}
}

func (r Rock) GetCoordinates() Coordinates {
	return r.Coordinates
}

func (r Rock) GetConsoleSprite() string {
	return "🪨"
}

func (t Tree) New(c Coordinates) Entity {
	return Tree{Coordinates: c}
}

func (t Tree) GetCoordinates() Coordinates {
	return t.Coordinates
}

func (t Tree) GetConsoleSprite() string {
	return "🌳"
}
