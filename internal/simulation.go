package internal

type Simulation struct {
	worldMap Map
}

func NewSimulation(worldMap Map) *Simulation {
	return &Simulation{
		worldMap: worldMap,
	}
}

func (s Simulation) Start() {
	initActions := []Action{}

	initActions = append(initActions, NewSpawnAction(Herbivore{}, 2))
	initActions = append(initActions, NewSpawnAction(Predator{}, 2))
	initActions = append(initActions, NewSpawnAction(Grass{}, 2))
	initActions = append(initActions, NewSpawnAction(Rock{}, 2))
	initActions = append(initActions, NewSpawnAction(Tree{}, 2))

	for _, ia := range initActions {
		ia.Perform(&s.worldMap)
	}

	Render(s.worldMap)
}
