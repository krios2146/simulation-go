package actions

import (
	"math/rand"

	"github.com/krios2146/simulation-go/internal"
	"github.com/krios2146/simulation-go/internal/entities"
	"github.com/krios2146/simulation-go/internal/pathfinder"
	"github.com/krios2146/simulation-go/internal/world"
)

type Action interface {
	Perform(m *world.Map)
}

type SpawnAction[T entities.Entity] struct {
	amount uint8
}

type TurnAction struct {
	pathFinder pathfinder.PathFinder
}

func NewSpawnAction[T entities.Entity](t T, amount uint8) *SpawnAction[T] {
	return &SpawnAction[T]{
		amount: amount,
	}
}

func (sa SpawnAction[T]) Perform(m *world.Map) {
	var x, y int
	var randCoordinates internal.Coordinates

	for i := 0; i < int(sa.amount); {
		x = rand.Intn(int(m.Width))
		y = rand.Intn(int(m.Height))

		randCoordinates = internal.Coordinates{X: uint8(x), Y: uint8(y)}

		if m.IsEmpty(randCoordinates) {
			entity := T.New(*new(T), randCoordinates)
			m.Add(entity, randCoordinates)
			i++
		}
	}
}

func NewTurnAction() *TurnAction {
	return &TurnAction{
		pathFinder: &pathfinder.BFSPathFinder{},
	}
}

func (ta TurnAction) Perform(m *world.Map) {
	creatures := world.FindByType(*new(entities.Creature), m)

	for _, c := range creatures {
		path := ta.pathFinder.FindPath(*m, c, c.InteractsWith())
		makeMove(path, m, c)
	}
}

func makeMove(path []internal.Coordinates, m *world.Map, creature entities.Creature) {
	panic("`makeMove` not implemented")
}
