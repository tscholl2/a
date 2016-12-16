package game

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"

	"github.com/tscholl2/a/entity"
	"github.com/tscholl2/a/utils"
)

type Game struct {
	Entities map[string]*entity.Entity // UUID -> Entity
	Turn     int                       // current turn number
	Size     int                       // maximum position in X or Y
	Stats    GameStats
}

func (g *Game) updateStats() {
	g.Stats.NumberOfAnimals = 0
	g.Stats.NumberOfPlants = 0
	animalSpecies := make(map[string]bool)
	plantSpecies := make(map[string]bool)
	for uuid := range g.Entities {
		if g.Entities[uuid].IsPlant {
			plantSpecies[g.Entities[uuid].Stats.SpeciesName] = true
			g.Stats.NumberOfPlants++
		} else {
			animalSpecies[g.Entities[uuid].Stats.SpeciesName] = true
			g.Stats.NumberOfAnimals++
		}
	}
	g.Stats.NumberOfAnimalSpecies = len(animalSpecies)
	g.Stats.NumberOfPlantSpecies = len(plantSpecies)
	log.Printf("GameStats: %+v\n", g.Stats)
}

func (g *Game) MakeWorld(worldSize int) {
	g.Size = worldSize
	g.Turn = 0
	g.Entities = make(map[string]*entity.Entity)
	for i := 0; i < 200; i++ {
		// Generate plants
		newEntity := generateRandomEntity(true)
		g.Entities[newEntity.UUID] = newEntity
	}
	for i := 0; i < 2; i++ {
		// Generate creature
		newEntity := generateRandomEntity(false)
		g.Entities[newEntity.UUID] = newEntity
	}

}

func generateRandomEntity(isPlant bool) *entity.Entity {
	var stats entity.Attributes
	stats.Type = utils.Types[rand.Intn(5)]
	stats.Defense = rand.Intn(20)
	stats.Strength = rand.Intn(20)
	stats.Endurance = rand.Intn(20)
	stats.Fortitude = rand.Intn(20)
	stats.Initiative = rand.Intn(20)
	if isPlant {
		stats.SpeciesName = fmt.Sprintf("%s plant %d", stats.Type, rand.Intn(10000))
	} else {
		stats.SpeciesName = fmt.Sprintf("%s creature %d", stats.Type, rand.Intn(10000))
	}
	stats.Priority.Attacker = rand.Intn(20)
	stats.Priority.Reproduction = rand.Intn(20)
	stats.Priority.Sleepy = rand.Intn(20)
	stats.Priority.Sleepy = rand.Intn(20)
	e := new(entity.Entity)
	e.Initialize(stats)
	e.IsPlant = isPlant
	return e
}

func (g *Game) Step() {
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
	// Add plants
	if math.Mod(float64(g.Turn), 5) == 0 && g.Stats.NumberOfPlants < 200 {
		newEntity := generateRandomEntity(true)
		g.Entities[newEntity.UUID] = newEntity
	}
	g.updateStats()
}

func (g *Game) getOrder() []*entity.Entity {
	entities := make([]*entity.Entity, len(g.Entities))
	entityInitative := make(map[string]int)
	for entityUUID := range g.Entities {
		entityInitative[entityUUID] = g.Entities[entityUUID].GetInitiative()
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
		if _, ok := g.Entities[e.UUID]; !ok {
			g.Entities[e.UUID] = e
		}
		if e.HP <= 0 {
			log.Printf("%s-%s died", e.Stats.SpeciesName, e.UUID)
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
