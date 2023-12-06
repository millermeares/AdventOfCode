package five

import (
	"fmt"
	"math"
)

type IngredientRange struct {
	beginning int
	length    int
}

// func (ir IngredientRange) end() int {
// 	return ir.beginning + ir.length
// }

// func (r Range) sourceEnd() int {
// 	return r.source + r.length
// }

// func (r Range) destEnd() int {
// 	return r.destination + r.length
// }

func getSeedRangesFromSeeds(seeds []int) []IngredientRange {
	var seedRanges []IngredientRange
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, IngredientRange{
			beginning: seeds[i],
			length:    seeds[i+1],
		})
	}
	return seedRanges
}

func lowestLocationFromRanges(a *Almanac, irs []IngredientRange) int {
	fannedOutIrs := a.fanoutIngredientRanges("seed", "location", irs)
	fmt.Println("Fanned out ingredient ranges:", fannedOutIrs)
	lowest := math.MaxInt
	for _, ir := range fannedOutIrs {
		if lowest > ir.beginning {
			lowest = ir.beginning
		}
	}
	return lowest
}

func (a *Almanac) fanoutIngredientRanges(source string, destination string, irs []IngredientRange) []IngredientRange {
	if source == destination {
		return irs
	}
	var newRanges []IngredientRange
	// rm := a.getRangeMapForSource(source)
	// for _, ir := range irs {

	// }
	return newRanges
}

// func (ir IngredientRange) fanoutIngredientRange(rm RangeMap) []IngredientRange {
// 	var irs []IngredientRange
// }
