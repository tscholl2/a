package game

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"time"

	"github.com/tscholl2/a/entity"
)

type Game struct {
	Entities map[string]*entity.Entity // UUID -> Entity
	Turn     int                       // current turn number
	Size     int                       // maximum position in X or Y
	Stats    GameStats
}

func (g *Game) Run() {
	ticker := time.NewTicker(time.Second).C
	select {
	case <-ticker:
		g.turn()
	}
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
	entities := make([]*entity.Entity, len(g.Entities))
	entityInitative := make(map[string]int)
	for entityUUID := range g.Entities {
		entityInitative[entityUUID] = g.Entities[entityUUID].Stats.Initiative
	}
	pairs := rankByInitiative(entityInitative)
	for i := range pairs {
		entities[i] = g.Entities[pairs[i].Key]
	}
	return entities
}

func rankByInitiative(entityInitative map[string]int) PairList {
	pl := make(PairList, len(entityInitative))
	i := 0
	for k, v := range entityInitative {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (g *Game) findAllInSquare(self *entity.Entity, x, y int) (things []*entity.Entity) {
	for _, e := range g.Entities {
		if int(math.Mod(float64(e.Position.X), float64(g.Size))) == int(math.Mod(float64(x), float64(g.Size))) && int(math.Mod(float64(e.Position.Y), float64(g.Size))) == int(math.Mod(float64(y), float64(g.Size))) && e.UUID != self.UUID {
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
