package internal

type Creature interface {
	Entity
	New(Coordinates) Entity
	InteractsWith() Entity
	MakeMove([]Coordinates, *Map)
}

type Predator struct {
	coordinates   Coordinates
	speed         uint8
	attackRating  uint8
	interactsWith Herbivore
}

type Herbivore struct {
	coordinates Coordinates
	speed       uint8
	health      uint8
}

func (p Predator) New(c Coordinates) Entity {
	return Predator{
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

func (p *Predator) MakeMove(path []Coordinates, m *Map) {
	if len(path) == 0 {
		return
	}
	if len(path) > int(p.speed) {
		moveTo := path[p.speed-1]
		m.Move(p.coordinates, moveTo)
		p.coordinates = moveTo
		return
	}

	var closest Coordinates
	if len(path) == 1 {
		closest = path[0]
	} else {
		closest = path[len(path)-2]
	}
	m.Move(p.coordinates, closest)
	p.coordinates = closest

	if herbivore, found := Find[Herbivore](*m, path[len(path)-1]); found {
		herbivore.health -= p.attackRating
	}
}

func (p Predator) GetConsoleSprite() string {
	return "🐺"
}

func (h Herbivore) New(c Coordinates) Entity {
	return Herbivore{
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

func (h Herbivore) GetConsoleSprite() string {
	return "🐑"
}

func (h *Herbivore) MakeMove(path []Coordinates, m *Map) {
	if len(path) == 0 {
		return
	}

	if len(path) > int(h.speed) {
		moveTo := path[h.speed-1]
		m.Move(h.coordinates, moveTo)
		h.coordinates = moveTo
		return
	}

	var closest Coordinates
	if len(path) == 1 {
		closest = path[0]
	} else {
		closest = path[len(path)-2]
	}
	m.Move(h.coordinates, closest)
	h.coordinates = closest

	m.Delete(path[len(path)-1])
	h.health += 25
}
