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
	total := 0
	total += e.Stats.Priority.Attacker
	total += e.Stats.Priority.Reproduction
	total += e.Stats.Priority.Sleepy
	total += e.Stats.Priority.Speed
	actionChoice := rand.Intn(total)
	if actionChoice < e.Stats.Priority.Attacker {
		return attack
	} else if actionChoice < e.Stats.Priority.Reproduction {
		return reproduce
	} else if actionChoice < e.Stats.Priority.Sleepy {
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
