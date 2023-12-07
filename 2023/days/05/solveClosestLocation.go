package five

import (
	"days"
	"math"
)

func GetDay() days.Day {
	return days.MakeDay(SolveClosestLocation, SolveClosestLocationWithRange, "05")
}

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

// the fundamental reality is that iterating over the seed inputs is going to be challenging.
// even something as simple as printing the seed input is CPU intensive.
// i need to either do this deterministically O(1) or *seriously* filter the input.
// the fanout solution ended up making it O(ranges input), rather than O(total numbers)
func SolveClosestLocationWithRange(input []string) int {
	a := parseAlmanac(input)
	seedRanges := getSeedRangesFromSeeds(a.seeds)
	return lowestLocationFromRanges(&a, seedRanges)
}
