package internal

import (
	"math/rand"
)

type Action interface {
	Perform(m *Map)
}

type SpawnAction[T Spawnable] struct {
	amount uint8
}

type Spawnable interface {
	New(Coordinates) Entity
}

type TurnAction struct{}

func NewSpawnAction[T Spawnable](amount uint8) *SpawnAction[T] {
	return &SpawnAction[T]{
		amount: amount,
	}
}

func (sa SpawnAction[T]) Perform(m *Map) {
	var x, y int
	var randCoordinates Coordinates

	for i := 0; i < int(sa.amount); {
		x = rand.Intn(int(m.Width))
		y = rand.Intn(int(m.Height))

		randCoordinates = Coordinates{X: uint8(x), Y: uint8(y)}

		if m.IsEmpty(randCoordinates) {
			entity := T.New(*new(T), randCoordinates)
			m.Add(entity, randCoordinates)
			i++
		}
	}
}

func NewTurnAction() *TurnAction {
	return &TurnAction{}
}

func (ta TurnAction) Perform(m *Map) {
	creatures := FindByType[Creature](*m)

	for _, c := range creatures {
		c.MakeMove(m)
	}
}
