package main

import (
	"github.com/krios2146/simulation-go/internal/actions"
	"github.com/krios2146/simulation-go/internal/entities"
	"github.com/krios2146/simulation-go/internal/ui"
	"github.com/krios2146/simulation-go/internal/world"
)

func main() {
	worldMap := world.NewMap(10, 10)

	initActions := []actions.Action{}

	initActions = append(initActions, new(actions.SpawnAction[entities.Herbivore]).New(2))
	initActions = append(initActions, new(actions.SpawnAction[entities.Predator]).New(2))
	initActions = append(initActions, new(actions.SpawnAction[entities.Grass]).New(2))
	initActions = append(initActions, new(actions.SpawnAction[entities.Rock]).New(2))
	initActions = append(initActions, new(actions.SpawnAction[entities.Tree]).New(2))

	for _, ia := range initActions {
		ia.Perform(worldMap)
	}

	ui.Render(worldMap)
}
