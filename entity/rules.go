package entity

import (
	"math"
	"math/rand"
)

const (
	attack    = "attack"
	reproduce = "reproduce"
	sleep     = "sleep"
	move      = "move"
)

func (e *Entity) chooseActionType(neighbors []*Entity) string {
	// Determine if possible to attack
	attackPossibility := 0
	for _, neighbor := range neighbors {
		if (neighbor.IsPlant && e.Stats.Vegetarian) || !e.Stats.Vegetarian {
			attackPossibility = e.Stats.Priority.Attacker
			break
		}
	}

	// Determine if its possible to reproduce
	reproducePossibility := 0
	if e.canReproduce() {
		reproducePossibility = e.Stats.Priority.Reproduction
	}

	// You can always sleep and move
	sleepPossibility := e.Stats.Priority.Sleepy
	movePossibility := e.Stats.Priority.Speed

	// Beacons trump all
	if e.BeaconPosition.X != -1 && e.BeaconPosition.Y != -1 {
		attackPossibility = 0
		reproducePossibility = 0
		sleepPossibility = 0
		movePossibility = 1
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

func (e *Entity) attackAction(targets []*Entity) []*Entity {
	// do stuff here
	return []*Entity{e, targets[0]}
}

func (e *Entity) reproduceAction() []*Entity {
	// do stuff here
	return []*Entity{e, new(Entity)}
}

func (e *Entity) sleepAction() []*Entity {
	// do stuff here
	return []*Entity{e}
}

func (e *Entity) moveAction() []*Entity {
	// Check if Beacon exists
	if e.BeaconPosition.X != -1 && e.BeaconPosition.Y != -1 {
		// Move toward beacon

		// Calculate unit vector
		v := Coordinate{e.BeaconPosition.X - e.Position.X, e.BeaconPosition.Y - e.Position.Y}
		vMag := math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
		vNorm := Coordinate{int(float64(v.X) / vMag), int(float64(v.Y) / vMag)}

		e.Position.X = e.Position.X + vNorm.X
		e.Position.Y = e.Position.Y + vNorm.Y
	} else {
		// Move randomly
		e.Position.X = e.Position.X + rand.Intn(3) - 1
		e.Position.Y = e.Position.Y + rand.Intn(3) - 1
	}
	// do stuff here
	return []*Entity{e}
}

func (e Entity) canReproduce() bool {
	return 0.75*float64(e.SP) > float64(e.MaxSP) && 0.75*float64(e.HP) > float64(e.MaxHP)
}
