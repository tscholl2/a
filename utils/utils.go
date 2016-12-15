package utils

// Types are the Wu-Xing for determining attacking
var Types = [...]string{"wood", "fire", "earth", "metal", "water"}

// CompareTypes returns whether i *generates* j
// or whether i *overcomes* j
func compareTypes(i, j int) (generates, overcomes bool) {
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

// CompareTypes returns whether i *generates* j
// or whether i *overcomes* j
func CompareType(i, j string) (generates, overcomes bool) {
	ii := 0
	jj := 0
	for t, theType := range Types {
		if theType == i {
			ii = t
		}
		if theType == j {
			jj = t
		}
	}
	return compareTypes(ii, jj)
}
