package actions

import (
	"math/rand"

	"github.com/krios2146/simulation-go/internal/coordinates"
	"github.com/krios2146/simulation-go/internal/entities"
	"github.com/krios2146/simulation-go/internal/world"
)

type spawnAction[T entities.Entity] interface {
	spawn(m *world.Map, amount uint8)
}

type SpawnAction[T entities.Entity] struct {
	amount uint8
}

func (sa SpawnAction[T]) New(amount uint8) *SpawnAction[T] {
	return &SpawnAction[T]{
		amount: amount,
	}
}

func (sa *SpawnAction[T]) Perform(m *world.Map) {
	sa.spawn(m)
}

func (sa *SpawnAction[T]) spawn(m *world.Map) {
	var x, y int
	var randCoordinates coordinates.Coordinates

	for i := 0; i < int(sa.amount); {
		x = rand.Intn(int(m.Width))
		y = rand.Intn(int(m.Height))

		randCoordinates = coordinates.Coordinates{X: uint8(x), Y: uint8(y)}

		if m.IsEmpty(randCoordinates) {
			entity := T.New(*new(T), randCoordinates)
			m.Add(entity, randCoordinates)
			i++
		}
	}
}
