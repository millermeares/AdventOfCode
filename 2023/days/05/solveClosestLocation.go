package five

import (
	"math"
)

func SolveClosestLocation(input []string) int {
	a := parseAlmanac(input)
	return a.getLowestLocation()
}

func (a Almanac) getLowestLocation() int {
	lowest := math.MaxInt
	for _, seed := range a.seeds {
		location := a.getMapping("seed", "location", seed)

		if location < lowest {
			lowest = location
		}
	}
	return lowest
}
