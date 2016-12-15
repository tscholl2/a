package game

import (
	"fmt"
	"math/rand"

	"github.com/tscholl2/a/entity"
	"github.com/tscholl2/a/utils"
)

func randomPlant() entity.Entity {
	var stats entity.Attributes
	stats.Type = utils.Types[rand.Intn(5)]
	stats.Defense = rand.Intn(20)
	stats.Strength = rand.Intn(20)
	stats.Endurance = rand.Intn(20)
	stats.Fortitude = rand.Intn(20)
	stats.Initiative = rand.Intn(20)
	stats.SpeciesName = fmt.Sprintf("%s plant %d", stats.Type, rand.Intn(10000))
	stats.Priority.Attacker = rand.Intn(20)
	stats.Priority.Reproduction = rand.Intn(20)
	stats.Priority.Sleepy = rand.Intn(20)
	stats.Priority.Sleepy = rand.Intn(20)
	var e entity.Entity
	e.Initialize(stats)
	e.IsPlant = true
	return e
}
