package internal

type Creature interface {
	Entity
	MakeMove(*Map)
}

type Predator struct {
	coordinates  Coordinates
	speed        uint8
	attackRating uint8
	pathFinder   PathFinder[Entity]
}

type Herbivore struct {
	coordinates Coordinates
	speed       uint8
	health      uint8
	pathFinder  PathFinder[Entity]
}

func (p Predator) GetCoordinates() Coordinates {
	return p.coordinates
}

func (p Predator) New(c Coordinates) Entity {
	return &Predator{
		coordinates:  c,
		speed:        2,
		attackRating: 25,
		pathFinder:   NewBFSPathFinder[*Herbivore](),
	}
}

func (p *Predator) MakeMove(m *Map) {
	path := p.pathFinder.FindPath(*m, p.coordinates)

	if len(path) == 0 {
		return
	}

	switch {
	case len(path) > int(p.speed):
		step := path[p.speed-1]
		p.move(step, m)

	case len(path) > 1:
		nearFood := path[len(path)-2]
		prey := path[len(path)-1]
		p.move(nearFood, m)
		p.attack(prey, m)

	default:
		prey := path[len(path)-1]
		p.attack(prey, m)
	}
}

func (p Predator) GetConsoleSprite() string {
	return "🐺"
}

func (p *Predator) move(c Coordinates, m *Map) {
	m.Move(p.coordinates, c)
	p.coordinates = c
}

func (p *Predator) attack(c Coordinates, m *Map) {
	if herbivore, found := Find[*Herbivore](*m, c); found {
		if p.attackRating > herbivore.health {
			herbivore.health = 0
		} else {
			herbivore.health -= p.attackRating
		}

		if herbivore.health <= 0 {
			m.Delete(c)
		}
	}
}

func (h Herbivore) GetCoordinates() Coordinates {
	return h.coordinates
}

func (h Herbivore) New(c Coordinates) Entity {
	return &Herbivore{
		coordinates: c,
		speed:       1,
		health:      100,
		pathFinder:  NewBFSPathFinder[Grass](),
	}
}

func (h Herbivore) GetConsoleSprite() string {
	return "🐑"
}

func (h *Herbivore) MakeMove(m *Map) {
	path := h.pathFinder.FindPath(*m, h.coordinates)

	if len(path) == 0 {
		return
	}

	switch {
	case len(path) > int(h.speed):
		step := path[h.speed-1]
		h.move(step, m)

	case len(path) > 1:
		nearFood := path[len(path)-2]
		food := path[len(path)-1]
		h.move(nearFood, m)
		h.eat(food, m)

	default:
		food := path[len(path)-1]
		h.eat(food, m)
	}
}

func (h *Herbivore) move(c Coordinates, m *Map) {
	m.Move(h.coordinates, c)
	h.coordinates = c
}

func (h *Herbivore) eat(c Coordinates, m *Map) {
	m.Delete(c)
	h.health += 25
}
