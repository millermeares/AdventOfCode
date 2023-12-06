package five

import (
	"fmt"
	"math"
)

type IngredientRange struct {
	beginning int
	length    int
}

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

func (a *Almanac) fanoutIngredientRanges(source string, destination string, currentIngredients []IngredientRange) []IngredientRange {
	if source == destination {
		return currentIngredients
	}
	rm := a.getRangeMapForSource(source)
	var irs []IngredientRange
	for _, curr := range currentIngredients {
		irs = append(irs, rm.fanoutIngredient(curr)...)
	}
	return irs // this is wrong, supposed to be recursive
	// return a.fanoutIngredientRanges(rm.to, destination, irs)
}

func (rm *RangeMap) fanoutIngredient(ir IngredientRange) []IngredientRange {
	var newRanges []Range
	// the ingredient range is in terms of "source". i just need to put it in terms of "destination"
	for _, r := range rm.ranges {
		if r.source+r.length < ir.beginning { // consider <=
			continue // ingredient range is after. go to next.
		}
		if ir.beginning+ir.length < r.source {
			continue // entire ingredient range is before range.
		}

		// same starting point.
		if ir.beginning == r.source {
			end := ir.length
			if ir.length > r.length {
				end = r.length // take smaller one
			}
			newRanges = append(newRanges, Range{
				source:      ir.beginning,
				destination: r.destination,
				length:      end,
			})
			continue
		}

		if ir.beginning > r.source {
			// r starts first
			rEnd := r.source + r.length
			irBegin := ir.beginning
			newRanges = append(newRanges, Range{
				source:      ir.beginning,
				destination: r.destination + (ir.beginning - r.source),
				length:      rEnd - irBegin,
			})
			continue
		}

		if ir.beginning < r.source {
			// ir starts first.
			irEnd := ir.beginning + ir.length
			rBegin := r.source
			newRanges = append(newRanges, Range{
				source:      ir.beginning,
				destination: r.destination,
				length:      irEnd - rBegin,
			})
			continue
		}
		panic("Did not match any cases, this is unexpected")
	}
	fmt.Println("New ranges:", newRanges, "from", ir, "and rangeMap", rm)
	return fillGaps(ir, newRanges)
}

// anything in the IngredientRange range must explicitly mapped to the same value.
func fillGaps(ir IngredientRange, fannedOutRanges []Range) []IngredientRange {
	var irs []IngredientRange
	ptr := ir.beginning
	for rIdx := 0; ptr < ir.beginning+ir.length && rIdx < len(fannedOutRanges); rIdx++ {
		// must increment ptr each iteration.
		r := fannedOutRanges[rIdx]
		// gap between ptr and r.source must be filled.
		if ptr < r.source {
			irs = append(irs, IngredientRange{
				beginning: ptr,
				length:    ptr - r.source,
			})
			// ptr = r.source not needed since we are copying destination next.
		}

		// not a gap - just copy mapped stuff over.
		irs = append(irs, IngredientRange{
			beginning: r.destination,
			length:    r.length,
		})
		ptr = r.source + r.length
	}
	if ptr < ir.beginning+ir.length {
		// fill end gap with a final ir.
		irs = append(irs, IngredientRange{
			beginning: ptr,
			length:    ir.beginning + ir.length - ptr,
		})
	}
	return irs
}
