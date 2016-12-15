package game

import (
	"fmt"
	"math/rand"

	"github.com/tscholl2/a/entity"
)

type Game struct {
	Entities map[string]*entity.Entity // UUID -> Entity
	Turn     int                       // current turn number
	Size     int                       // maximum position in X or Y
	Stats    GameStats
}

func (g *Game) turn() {
	g.Turn++
	entities := g.getOrder()
	for _, e := range entities {
		if _, ok := g.Entities[e.UUID]; !ok {
			continue
		}
		neighbors := g.findAllInSquare(e, e.Position.X, e.Position.Y)
		action := e.GetAction(neighbors)
		// TODO save some info about action
		updates := action.Perform()
		g.mergeUpdates(updates)
	}
}

func (g *Game) getOrder() []*entity.Entity {
	// TODO
	return nil
}

func (g *Game) findAllInSquare(self *entity.Entity, x, y int) (things []*entity.Entity) {
	for _, e := range g.Entities {
		if e.Position.X == x && e.Position.Y == y && e.UUID != self.UUID {
			things = append(things, e)
		}
	}
	return
}

func (g *Game) mergeUpdates(updates []*entity.Entity) {
	for _, e := range updates {
		if e.UUID == "" {
			e.UUID = fmt.Sprintf("%d", rand.Int63())
			g.Entities[e.UUID] = e
		}
		if e.HP == 0 {
			delete(g.Entities, e.UUID)
		}
	}
	return
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
