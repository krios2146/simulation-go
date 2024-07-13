package internal

type Creature interface {
	New(c Coordinates) Entity
	InteractsWith() Entity
	GetCoordinates() Coordinates
}

type Predator struct {
	Coordinates  Coordinates
	Speed        uint8
	AttackRating uint8
}

type Herbivore struct {
	Coordinates Coordinates
	Speed       uint8
	Health      uint8
}

func (p Predator) New(c Coordinates) Entity {
	return &Predator{
		Coordinates:  c,
		Speed:        2,
		AttackRating: 50,
	}
}

func (p Predator) InteractsWith() Entity {
	return Herbivore{}
}

func (p Predator) GetCoordinates() Coordinates {
	return p.Coordinates
}

func (p Predator) GetConsoleSprite() rune {
	// 🐺
	return '\U0001F43A'
}

func (h Herbivore) New(c Coordinates) Entity {
	return &Herbivore{
		Coordinates: c,
		Speed:       1,
		Health:      100,
	}
}

func (h Herbivore) InteractsWith() Entity {
	return Grass{}
}

func (h Herbivore) GetCoordinates() Coordinates {
	return h.Coordinates
}

func (h Herbivore) GetConsoleSprite() rune {
	// 🐑
	return '\U0001F411'
}
