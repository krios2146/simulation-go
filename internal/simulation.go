package internal

import (
	"fmt"
	"time"
)

type Simulation struct {
	worldMap Map
}

func NewSimulation(worldMap Map) *Simulation {
	return &Simulation{
		worldMap: worldMap,
	}
}

func (s Simulation) Start() {
	for _, ia := range initActions() {
		ia.Perform(&s.worldMap)
	}
	Render(s.worldMap)

	turnAction := NewTurnAction()

	for !isOver(s.worldMap) {
		turnAction.Perform(&s.worldMap)
		Render(s.worldMap)
		time.Sleep(time.Second)
	}
}

func initActions() []Action {
	initActions := []Action{}

	initActions = append(initActions, NewSpawnAction(Herbivore{}, 1))
	initActions = append(initActions, NewSpawnAction(Predator{}, 0))
	initActions = append(initActions, NewSpawnAction(Grass{}, 1))
	initActions = append(initActions, NewSpawnAction(Rock{}, 0))
	initActions = append(initActions, NewSpawnAction(Tree{}, 0))

	return initActions
}

func isOver(m Map) bool {
	if len(FindByType[Herbivore](m)) == 0 {
		return true
	} else {
		fmt.Printf("[isOver]: found more than zero Herbivore with `FindByType` call\n")
	}
	if len(FindByType[Grass](m)) == 0 {
		return true
	}
	return false
}
