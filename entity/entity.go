package entity

import "log"

// Entity is a players of the game
type Entity struct {
	Stats          Attributes // List of stats, set by user-created entities
	UUID           string     // UUID is unique for all entities
	IsPlant        bool       // determine if its a plant (users cannot create plants)
	Age            int        // age, increasesd every step
	Generation     int        // generation, increased upon reproduction
	Position       Coordinate // position of entity
	BeaconPosition Coordinate // position of species-specific beacon
	MaxHP          int        // maximum Hit Points
	HP             int        // current Hit Points
	MaxSP          int        // maximum Stamina Points
	SP             int        // current Stamina Points
	History        []string   // History of actions
}

// Action is something to be preformed by the game
type Action struct {
	Description string
	Perform     func() []*Entity
}

// GetAction takes a list of targets and returns an action this entity wants to do.
func (e *Entity) GetAction(neighbors []*Entity) Action {
	t := e.chooseActionType(neighbors)
	e.History = append(e.History, t)
	e.Age++
	if t != donothing {
		log.Printf("%s-%s (%d/%d, %d/%d) is performing %s", e.Stats.SpeciesName, e.UUID, e.HP, e.MaxHP, e.SP, e.MaxSP, t)
	}
	switch t {
	case "attack":
		return Action{
			Description: "attack",
			Perform: func() (updates []*Entity) {
				return e.attackAction(neighbors)
			}}
	case "sleep":
		return Action{
			Description: "sleep",
			Perform: func() (updates []*Entity) {
				return e.sleepAction()
			}}
	case "move":
		return Action{
			Description: "move",
			Perform: func() (updates []*Entity) {
				return e.moveAction()
			}}
	case "reproduce":
		return Action{
			Description: "reproduce",
			Perform: func() (updates []*Entity) {
				return e.reproduceAction()
			}}
	default:
		return Action{
			Description: "",
			Perform:     func() (updates []*Entity) { return nil },
		}
	}

}

// Coordinate keeps track of position
type Coordinate struct {
	X, Y int
}

// Attributes are things the user sets for their creature
type Attributes struct {
	SpeciesName string      `json:"species_name"` // user determines Species name
	Type        string      `json:"type"`         // wood, fire, earth, metal, water
	Initiative  int         `json:"initiative"`   // affects ordering in game
	Strength    int         `json:"strength"`     // amount hitpoints absorbed, x2 for advantage
	Defense     int         `json:"defense"`      // amount to decrease hitpoints absorbed, x2 for advantage
	Endurance   int         `json:"endurance"`    // decreases stamina used for moving
	Fortitude   int         `json:"fortitude"`    // amount of stamina gained for sleeping
	Priority    ActionStats `json:"priority"`     // Priority determines what actions will be chosen
	Vegetarian  bool        `json:"vegetarian"`   // when attacking, chooses plants as target
	Aggressive  bool        `json:"aggressive"`   // when attacking, chooses lower defense as target
	Scavenger   bool        `json:"scavenger"`    // when attacking, chooses lower HP as target
}

// ActionStats are the priorities for the action.
// They are just arbitrary numbers.
// These priorities are normalized when determing which action to undertake
// to determine the priority with probabilities.
type ActionStats struct {
	Attacker int `json:"attacker"` // increases probability for attack action
	// the decision on which Entity to attack is chosen by other attributes
	Speed        int `json:"speed"`        // increases probability for move action
	Reproduction int `json:"reproduction"` // increases probability for reproducing
	Sleepy       int `json:"sleepy"`       // increases probability for sleep action
}
