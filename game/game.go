package game

import (
	"github.com/tscholl2/a/entity"
)

type Game struct {
	Entities     map[string]entity.Entity // UUID -> Entity
	OrderOfTurns []string                 // list of Entity UUIDs
	Turn         int                      // current turn number
	Size         int                      // maximum position in X or Y
}

// functions for learning about game

// NumberOfPlants returns number of plants

// NumberOfPlantSpecies returns number of plant species
