package game

import (
	"github.com/tscholl2/a/entity"
)

type Game struct {
	Entities     map[string]entity.Entity // UUID -> Entity
	OrderOfTurns []string                 // list of Entity UUIDs
	Turn         int                      // current turn number
	Size         int                      // maximum position in X or Y
	Stats        GameStats
}

// GameStats are useful metrics for the current game state
// which have to be parsed by looking at the entities
type GameStats struct {
	NumberOfPlants        int
	NumberOfPlantSpecies  int
	NumberOfAnimals       int
	NumberOfAnimalSpecies int
	TypeCount             [5]int
}
