package utils

var Types = [...]string{"wood", "fire", "earth", "metal", "water"}

// CompareTypes returns whether i *generates* j
// or whether i *overcomes* j
func CompareTypes(i, j int) (generates, overcomes bool) {
	foo := i + 1
	if foo > 4 {
		foo = foo - 5
	}
	if foo == j {
		generates = true
		return
	}
	foo = i + 2
	if foo > 4 {
		foo = foo - 5
	}
	if foo == j {
		overcomes = true
	}
	return
}
