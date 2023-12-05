package five

type SeedRange struct {
	beginning int
	length    int
}

func getSeedRangesFromSeeds(seeds []int) []SeedRange {
	var seedRanges []SeedRange
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			beginning: seeds[i],
			length:    seeds[i+1],
		})
	}
	return seedRanges
}

// func (a *Almanac) lowestLocationFromRange(sr SeedRange) int {
// 	lowest := math.MaxInt
// 	for i := sr.beginning; i < sr.beginning+sr.length; i++ {
// 		location := a.getMapping("seed", "location", i)
// 		if lowest > location {
// 			lowest = location
// 		}
// 	}
// 	return lowest
// }
