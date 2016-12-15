package entity

import "math/rand"

const (
	attack    = "attack"
	reproduce = "reproduce"
	sleep     = "sleep"
	move      = "move"
)

func (e *Entity) chooseActionType(neighbors []*Entity) string {
	types := []string{attack, reproduce, sleep, move}
	return types[rand.Intn(len(types))]
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
	// do stuff here
	return []*Entity{e}
}
