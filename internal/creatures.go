package internal

type Creature interface {
	New(c Coordinates) Entity
	InteractsWith() Entity
	GetCoordinates() Coordinates
}

type Predator struct {
	coordinates  Coordinates
	speed        uint8
	attackRating uint8
}

type Herbivore struct {
	coordinates Coordinates
	speed       uint8
	health      uint8
}

func (p Predator) New(c Coordinates) Entity {
	return &Predator{
		coordinates:  c,
		speed:        2,
		attackRating: 50,
	}
}

func (p Predator) InteractsWith() Entity {
	return Herbivore{}
}

func (p Predator) GetCoordinates() Coordinates {
	return p.coordinates
}

func (p Predator) GetConsoleSprite() rune {
	// 🐺
	return '\U0001F43A'
}

func (h Herbivore) New(c Coordinates) Entity {
	return &Herbivore{
		coordinates: c,
		speed:       1,
		health:      100,
	}
}

func (h Herbivore) InteractsWith() Entity {
	return Grass{}
}

func (h Herbivore) GetCoordinates() Coordinates {
	return h.coordinates
}

func (h Herbivore) GetConsoleSprite() rune {
	// 🐑
	return '\U0001F411'
}
