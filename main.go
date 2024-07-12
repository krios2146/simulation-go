package main

import (
	"github.com/krios2146/simulation-go/internal/coordinates"
	"github.com/krios2146/simulation-go/internal/entities"
	"github.com/krios2146/simulation-go/internal/ui"
	"github.com/krios2146/simulation-go/internal/world"
)

func main() {
	worldMap := world.NewMap(10, 10)
	var entity entities.Entity
	var coordinates coordinates.Coordinates

	coordinates.X = 1
	coordinates.Y = 1
	entity = entities.Rock{Coordinates: coordinates}

	worldMap.Add(entity, coordinates)

	coordinates.X = 2
	coordinates.Y = 2
	entity = entities.Grass{Coordinates: coordinates}

	worldMap.Add(entity, coordinates)

	coordinates.X = 8
	coordinates.Y = 6
	entity = entities.Herbivore{Coordinates: coordinates}

	worldMap.Add(entity, coordinates)

	coordinates.X = 6
	coordinates.Y = 8
	entity = entities.Predator{Coordinates: coordinates}

	worldMap.Add(entity, coordinates)

	ui.Render(*worldMap)
}
