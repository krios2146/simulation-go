package world

import (
	"github.com/krios2146/simulation-go/internal/coordinates"
	"github.com/krios2146/simulation-go/internal/entities"
)

type Map struct {
	values map[coordinates.Coordinates]entities.Entity
	Height uint8
	Width  uint8
}

func NewMap(height uint8, width uint8) *Map {
	return &Map{
		values: map[coordinates.Coordinates]entities.Entity{},
		Height: height,
		Width:  width,
	}
}

func (m *Map) Find(coordinates coordinates.Coordinates) (entity entities.Entity, found bool) {
	entity, found = m.values[coordinates]
	return
}

func (m *Map) FindCreatures() []entities.Creature {
	creatures := []entities.Creature{}

	for _, entity := range m.values {
		if entity.CanMove() {
			creatures = append(creatures, entity.(entities.Creature))
		}
	}

	return creatures
}

func (m *Map) FindHerbivores() []entities.Herbivore {
	herbs := []entities.Herbivore{}

	for _, v := range m.values {
		if h, ok := v.(entities.Herbivore); ok {
			herbs = append(herbs, h)
		}
	}

	return herbs
}

func (m *Map) FindGrass() []entities.Grass {
	herbs := []entities.Grass{}

	for _, v := range m.values {
		if h, ok := v.(entities.Grass); ok {
			herbs = append(herbs, h)
		}
	}

	return herbs
}

func (m *Map) Add(entity entities.Entity, coordinates coordinates.Coordinates) {
	m.values[coordinates] = entity
}
