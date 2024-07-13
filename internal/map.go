package internal

type Map struct {
	values map[Coordinates]Entity
	Height uint8
	Width  uint8
}

func NewMap(height uint8, width uint8) *Map {
	return &Map{
		values: map[Coordinates]Entity{},
		Height: height,
		Width:  width,
	}
}

func (m Map) IsEmpty(c Coordinates) bool {
	_, occupied := m.values[c]
	return !occupied
}

func (m Map) Find(coordinates Coordinates) (entity Entity, found bool) {
	entity, found = m.values[coordinates]
	return
}

func FindByType[T Entity](e T, m Map) []T {
	entities := []T{}

	for _, v := range m.values {
		if e, ok := v.(T); ok {
			entities = append(entities, e)
		}
	}

	return entities
}

func (m *Map) Add(entity Entity, coordinates Coordinates) {
	m.values[coordinates] = entity
}
