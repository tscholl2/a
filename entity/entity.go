package entity

// Entities are the players of the game
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

// Coordinate keeps track of position
type Coordinate struct {
	x, y int
}

// Attributes are things the user sets for their creature
type Attributes struct {
	SpeciesName string      // user determines Species name
	Type        string      // wood, fire, earth, metal, water
	Initiative  int         // affects ordering in game
	Strength    int         // amount hitpoints absorbed, x2 for advantage
	Defense     int         // amount to decrease hitpoints absorbed, x2 for advantage
	Endurance   int         // decreases stamina used for moving
	Fortitude   int         // amount of stamina gained for sleeping
	Priority    ActionStats // Priority determines what actions will be chosen
	Vegetarian  bool        // when attacking, chooses plants as target
	Aggressive  bool        // when attacking, chooses lower defense as target
	Scavenger   bool        // when attacking, chooses lower HP as target
}

// ActionStats are the priorities for the action
// these priorities are normalized when determing which action to undertake
// to determine the priority with probabilities
type ActionStats struct {
	Attacker int // increases probability for attack action
	// the decision on which Entity to attack is chosen by other attributes
	Speed        int // increases probability for move action
	Reproduction int // increases probability for reproducing
	Sleepy       int // increases probability for sleep action
}
