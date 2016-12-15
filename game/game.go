package game

import (
	"github.com/tscholl2/a/entity"
)

type Game struct {
	Entities     map[string]entity.Entity // UUID -> Entity
	OrderOfTurns []string                 // list of Entity UUIDs
}
