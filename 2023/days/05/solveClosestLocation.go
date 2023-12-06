package five

import (
	"math"
)

func SolveClosestLocation(input []string) int {
	a := parseAlmanac(input)
	return a.getLowestLocation()
}

// part 1
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
// there is a finite amount of ranges. map ranges to each other? then iterate over those ranges searching for matches?
// input: range. output: []range that returns the ranges that those maps are split into. not too crazy.
// then for each resultRange in []range, do the same operation.
func SolveClosestLocationWithRange(input []string) int {
	a := parseAlmanac(input)
	seedRanges := getSeedRangesFromSeeds(a.seeds)
	return lowestLocationFromRanges(&a, seedRanges)
}
