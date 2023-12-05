package five

import (
	"fmt"
	"math"
)

func SolveClosestLocation(input []string) int {
	a := parseAlmanac(input)
	return a.getLowestLocation()
}

func SolveClosestLocationWithRange_BruteForce(input []string) int {
	a := parseAlmanac(input)
	seedRanges := getSeedRangesFromSeeds(a.seeds)
	lowest := math.MaxInt
	selectionSeedMap := map[int]int{}
	for rangeNum, sr := range seedRanges {
		fmt.Println("Evaluating range num", rangeNum)
		for seed := sr.beginning; seed < sr.beginning+sr.length; seed++ {
			location := getLocationOfSeed(&a, seed, selectionSeedMap)
			if lowest > location {
				lowest = location
			}
		}
	}
	return lowest
}

func getLocationOfSeed(a *Almanac, seed int, seedLocationMap map[int]int) int {
	val, exists := seedLocationMap[seed]
	if exists {
		return val
	}
	val = a.getMapping("seed", "location", seed)
	seedLocationMap[seed] = val
	return val

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
