package world

import (
	"github.com/krios2146/simulation-go/internal"
	"github.com/krios2146/simulation-go/internal/entities"
)

type Map struct {
	values map[internal.Coordinates]entities.Entity
	Height uint8
	Width  uint8
}

func NewMap(height uint8, width uint8) *Map {
	return &Map{
		values: map[internal.Coordinates]entities.Entity{},
		Height: height,
		Width:  width,
	}
}

func (m Map) IsEmpty(c internal.Coordinates) bool {
	_, occupied := m.values[c]
	return !occupied
}

func (m Map) Find(coordinates internal.Coordinates) (entity entities.Entity, found bool) {
	entity, found = m.values[coordinates]
	return
}

func FindByType[T entities.Entity](e T, m Map) []T {
	entities := []T{}

	for _, v := range m.values {
		if e, ok := v.(T); ok {
			entities = append(entities, e)
		}
	}

	return entities
}

func (m *Map) Add(entity entities.Entity, coordinates internal.Coordinates) {
	m.values[coordinates] = entity
}
