package five

import (
	"math"
)

type IngredientRange struct {
	beginning int
	length    int
}

func (ir IngredientRange) end() int {
	return ir.beginning + ir.length
}

func (r Range) sourceEnd() int {
	return r.source + r.length
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
	rm := a.getRangeMapForSource(source)
	for _, ir := range irs {
		newRanges = append(newRanges, rm.fanoutIngredientRange(ir)...)
	}
	return a.fanoutIngredientRanges(rm.to, destination, newRanges)
}

// the ingredient range is current in the terms of rm.to and [r.source]. need to turn it into rm.from and [r.destination]
// ir.length should be equal to the sum of ranges.length
func (rm RangeMap) fanoutIngredientRange(ir IngredientRange) []IngredientRange {
	var irs []IngredientRange

	mappingPointer := ir.beginning

	for _, r := range rm.ranges {
		if mappingPointer == ir.end() {
			break // all mapped!
		}

		if !r.inMappingRange(mappingPointer) && mappingPointer < r.source {
			// this was added to address ir.beginning less than min([r.source])
			toSelfEnd := int(math.Min(float64(r.source), float64(ir.end())))
			irs = append(irs, IngredientRange{
				beginning: mappingPointer,
				length:    toSelfEnd - mappingPointer,
			})
			mappingPointer = toSelfEnd
		}

		if r.inMappingRange(mappingPointer) {
			// from mappingPointer to r.source needs to map to same value.
			if mappingPointer < r.source {
				irs = append(irs, IngredientRange{
					beginning: mappingPointer,
					length:    r.source - mappingPointer,
				})

				mappingPointer = r.source
			}

			// map from mappingPointer to min(r.sourceEnd(), ir.end())) to r.destination
			rangeEnd := int(math.Min(float64(r.sourceEnd()), float64(ir.end())))
			irs = append(irs, IngredientRange{
				beginning: r.getMapping(mappingPointer),
				length:    rangeEnd - mappingPointer,
			})
			mappingPointer = rangeEnd
		}
	}

	// this handles 2 cases - 1. there no ranges in mapping range and 2. part at the end of the ir is not mapped to anything.
	if mappingPointer != ir.end() {
		irs = append(irs, IngredientRange{
			beginning: mappingPointer,
			length:    ir.end() - mappingPointer,
		})
	}
	return irs
}
