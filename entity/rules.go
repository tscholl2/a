package entity

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/tscholl2/a/utils"
)

const (
	attack    = "attack"
	reproduce = "reproduce"
	sleep     = "sleep"
	move      = "move"
	donothing = "donothing"
)

// Initialize creates a new Entity based on the stats supplied
func (e *Entity) Initialize(stats Attributes) {
	// Normalize stats
	if stats.Defense < 0 {
		stats.Defense = 0
	}
	if stats.Endurance < 0 {
		stats.Defense = 0
	}
	if stats.Fortitude < 0 {
		stats.Defense = 0
	}
	if stats.Initiative < 0 {
		stats.Defense = 0
	}
	if stats.Strength < 0 {
		stats.Defense = 0
	}
	total := stats.Defense + stats.Endurance + stats.Fortitude + stats.Initiative + stats.Strength + 5
	stats.Defense = 1 + int(20*float64(stats.Defense+1)/float64(total))
	stats.Endurance = 1 + int(20*float64(stats.Endurance+1)/float64(total))
	stats.Fortitude = 1 + int(20*float64(stats.Fortitude+1)/float64(total))
	stats.Initiative = 1 + int(20*float64(stats.Initiative+1)/float64(total))
	stats.Strength = 1 + int(20*float64(stats.Strength+1)/float64(total))

	// Normalize priorities
	if stats.Priority.Attacker < 0 {
		stats.Priority.Attacker = 0
	}
	if stats.Priority.Reproduction < 0 {
		stats.Priority.Reproduction = 0
	}
	if stats.Priority.Sleepy < 0 {
		stats.Priority.Sleepy = 0
	}
	if stats.Priority.Speed < 0 {
		stats.Priority.Speed = 0
	}
	total = stats.Priority.Attacker + stats.Priority.Reproduction + stats.Priority.Sleepy + stats.Priority.Speed + 4
	stats.Priority.Attacker = 1 + int(20*float64(stats.Priority.Attacker+1)/float64(total))
	stats.Priority.Reproduction = 1 + int(20*float64(stats.Priority.Reproduction+1)/float64(total))
	stats.Priority.Sleepy = 1 + int(20*float64(stats.Priority.Sleepy+1)/float64(total))
	stats.Priority.Speed = 1 + int(20*float64(stats.Priority.Speed+1)/float64(total))

	e.Stats = stats
	e.UUID = fmt.Sprintf("%d", rand.Int63())
	e.Age = 0
	e.Generation = 0
	e.Position.X = rand.Int()
	e.Position.Y = rand.Int()
	e.MaxHP = 10 + rand.Intn(e.Stats.Fortitude) + rand.Intn(e.Stats.Fortitude)
	e.HP = e.MaxHP
	e.MaxSP = 100 + rand.Intn(e.Stats.Endurance) + rand.Intn(e.Stats.Endurance)
	e.SP = e.MaxSP
	e.BeaconPosition.X = -1
	e.BeaconPosition.Y = -1
	log.Printf("Created %s-%s with %d HP, stats:%v", e.Stats.SpeciesName, e.UUID, e.HP, stats)
}

func (e Entity) canReproduce() bool {
	return float64(e.SP) > 0.75*float64(e.MaxSP) && float64(e.HP) > 0.75*float64(e.MaxHP)
}

func (e Entity) canAttack() bool {
	return e.SP > 0
}

func (e Entity) canMove() bool {
	return e.SP > 0
}

func (e *Entity) chooseActionType(neighbors []*Entity) string {
	// Plants do nothing
	if e.IsPlant {
		return donothing
	}

	attackPossibility := 0
	reproducePossibility := 0
	movePossibility := 0
	sleepPossibility := e.Stats.Priority.Sleepy // You can always sleep

	// Determine if possible to attack
	if e.canAttack() {
		for _, neighbor := range neighbors {
			if (neighbor.IsPlant && e.Stats.Vegetarian) || !e.Stats.Vegetarian {
				attackPossibility = e.Stats.Priority.Attacker
				break
			}
		}
	}

	// Determine if its possible to reproduce
	if e.canReproduce() {
		reproducePossibility = e.Stats.Priority.Reproduction
	}

	if e.canMove() {
		movePossibility = e.Stats.Priority.Speed
		// Beacons trump all
		if e.BeaconPosition.X != -1 && e.BeaconPosition.Y != -1 {
			attackPossibility = 0
			reproducePossibility = 0
			sleepPossibility = 0
			movePossibility = 1
		}
	}

	actionChoice := rand.Intn(attackPossibility + reproducePossibility + sleepPossibility + movePossibility)
	if actionChoice < attackPossibility {
		return attack
	} else if actionChoice < reproducePossibility {
		return reproduce
	} else if actionChoice < sleepPossibility {
		return sleep
	} else {
		return move
	}
}

func (e *Entity) hasPlantAdvantageAgainst(e2 *Entity) bool {
	_, generate := utils.CompareType(e2.Stats.Type, e.Stats.Type)
	return generate
}

func (e *Entity) hasAttackAdvantageAgainst(e2 *Entity) bool {
	overcome, _ := utils.CompareType(e.Stats.Type, e2.Stats.Type)
	return overcome
}

func (e *Entity) hasAttackDisadvantageAgainst(e2 *Entity) bool {
	overcome, _ := utils.CompareType(e2.Stats.Type, e.Stats.Type)
	return overcome
}

func (e *Entity) attackAction(targets []*Entity) []*Entity {
	// Choose target
	lowestDefense := 0
	lowestHP := 0
	targetID := 0
	for i, target := range targets {
		if e.Stats.Vegetarian && target.IsPlant {
			if e.hasPlantAdvantageAgainst(target) {
				targetID = i
				break
			}
			if target.Stats.Defense < lowestDefense && e.Stats.Aggressive && !e.hasAttackDisadvantageAgainst(target) {
				lowestDefense = target.Stats.Defense
				targetID = i
			}
			if target.HP < lowestHP && e.Stats.Scavenger && !e.hasAttackDisadvantageAgainst(target) {
				lowestHP = target.HP
				targetID = i
			}
		} else {
			if e.hasAttackAdvantageAgainst(target) {
				targetID = i
				break
			}
			if target.Stats.Defense < lowestDefense && e.Stats.Aggressive && !e.hasAttackDisadvantageAgainst(target) {
				lowestDefense = target.Stats.Defense
				targetID = i
			}
			if target.HP < lowestHP && e.Stats.Scavenger && !e.hasAttackDisadvantageAgainst(target) {
				lowestHP = target.HP
				targetID = i
			}
		}
	}

	// Attack target
	target := targets[targetID]
	hpAttack := rand.Intn(10) + rand.Intn(e.Stats.Strength)
	if e.hasAttackAdvantageAgainst(target) {
		log.Println("Has advantage!")
		hpAttack += rand.Intn(10) + rand.Intn(e.Stats.Strength)
	}
	hpDefense := rand.Intn(10) + rand.Intn(target.Stats.Defense)
	if e.hasAttackDisadvantageAgainst(target) {
		log.Println("Has disadvantage!")
		hpDefense += rand.Intn(10) + rand.Intn(target.Stats.Defense)
	}
	log.Println(hpAttack, hpDefense)
	hpTotal := hpAttack - hpDefense
	if hpTotal > target.HP {
		hpTotal = target.HP
	}
	if hpTotal < 0 {
		hpTotal = 0
	}
	e.SP = e.SP - 10
	e.HP = e.HP + hpTotal
	target.HP = target.HP - hpTotal
	log.Printf("%s-%s attacked %s-%s for %d damage", e.Stats.SpeciesName, e.UUID, target.Stats.SpeciesName, target.UUID, hpTotal)

	return []*Entity{e, target}
}

func (e *Entity) reproduceAction() []*Entity {
	newStats := e.Stats
	for _, action := range e.History {
		switch action {
		case attack:
			newStats.Priority.Attacker++
		case reproduce:
			newStats.Priority.Reproduction++
		case sleep:
			newStats.Priority.Sleepy++
		case move:
			newStats.Priority.Speed++
		}
	}
	// create new entity
	newEntity := new(Entity)
	newEntity.Initialize(newStats)
	newEntity.Generation = e.Generation + 1
	log.Printf("%s-%s created offspring: %s-%s", e.Stats.SpeciesName, e.UUID, newEntity.Stats.SpeciesName, newEntity.UUID)

	e.HP = e.HP / 2
	e.SP = e.SP / 2
	return []*Entity{e, newEntity}
}

func (e *Entity) sleepAction() []*Entity {
	e.SP += rand.Intn(e.Stats.Fortitude)
	if e.SP > e.MaxSP {
		e.SP = e.MaxSP
	}
	log.Printf("%s-%s sleeping", e.Stats.SpeciesName, e.UUID)
	return []*Entity{e}
}

func (e *Entity) moveAction() []*Entity {
	if e.BeaconPosition.X != -1 && e.BeaconPosition.Y != -1 {
		// If Beacon exists, move toward beacon

		// Calculate unit vector
		v := Coordinate{e.BeaconPosition.X - e.Position.X, e.BeaconPosition.Y - e.Position.Y}
		vMag := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
		vNorm := Coordinate{int(float64(v.X) / vMag), int(float64(v.Y) / vMag)}

		e.Position.X = e.Position.X + vNorm.X
		e.Position.Y = e.Position.Y + vNorm.Y

	} else {
		// Move randomly
		log.Println("Moving randomly")
		fmt.Println(rand.Intn(3) - 1)
		e.Position.X = e.Position.X + rand.Intn(3) - 1
		e.Position.Y = e.Position.Y + rand.Intn(3) - 1
	}
	log.Printf("%s-%s moved to %d,%d", e.Stats.SpeciesName, e.UUID, e.Position.X, e.Position.Y)
	return []*Entity{e}
}
