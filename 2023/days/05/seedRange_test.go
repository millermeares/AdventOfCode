package five

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestWithSeedRange(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	location := SolveClosestLocationWithRange(input)
	assert.Equal(t, 46, location)
}

func TestFanout_SmallestBefore(t *testing.T) {
	input := []string{
		"seeds: 48 52",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}
	a := parseAlmanac(input)
	srs := getSeedRangesFromSeeds(a.seeds) // [48, 52]
	fannedOut := a.fanoutIngredientRanges("seed", "soil", srs)
	fmt.Println(fannedOut)
	// beginning = 48, length = 52. so source is effective from 48 to 100.
	// so in terms of soil, this would be broken up into 3 categories.
	// 48 to 49 map to 48 and 49.
	// 50 to 98 map to 52 to 100.
	// 98 to 99 map to 50 and 51.

	// {50 2}, {48 2}, {52 48}
	// they will be in order of source.
	assert.Equal(t, srs[0].length, lengthSum(fannedOut))

	irOne := fannedOut[0]
	assert.Equal(t, 48, irOne.beginning)
	assert.Equal(t, 2, irOne.length)

	irTwo := fannedOut[1]
	assert.Equal(t, 52, irTwo.beginning)
	assert.Equal(t, 48, irTwo.length)

	irThree := fannedOut[2]
	assert.Equal(t, 50, irThree.beginning)
	assert.Equal(t, 2, irThree.length)
}

func TestFanout_NoMatchEnd(t *testing.T) {
	input := []string{
		"seeds: 110 52",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}
	a := parseAlmanac(input)
	srs := getSeedRangesFromSeeds(a.seeds) // [110, 52]
	fannedOut := a.fanoutIngredientRanges("seed", "soil", srs)
	fmt.Println(fannedOut)
	assert.Equal(t, 1, len(fannedOut))
	assert.Equal(t, srs[0], fannedOut[0])
}

func TestFanout_NoMatchBeginning(t *testing.T) {
	input := []string{
		"seeds: 30 5",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}
	a := parseAlmanac(input)
	srs := getSeedRangesFromSeeds(a.seeds)
	fannedOut := a.fanoutIngredientRanges("seed", "soil", srs)
	fmt.Println(fannedOut)
	assert.Equal(t, 1, len(fannedOut))
	assert.Equal(t, srs[0], fannedOut[0])
}

func TestFanout_NoMatchMiddle(t *testing.T) {
	input := []string{
		"seeds: 25 3", // 25-8 map to 25-8
		"",
		"seed-to-soil map:",
		"20 10 10", // 10-20 map to 20-30
		"40 30 11", // 30-41 map to 40-51
		"",
	}
	a := parseAlmanac(input)
	srs := getSeedRangesFromSeeds(a.seeds)
	fannedOut := a.fanoutIngredientRanges("seed", "soil", srs)
	fmt.Println(fannedOut)
	assert.Equal(t, 1, len(fannedOut))
	assert.Equal(t, srs[0], fannedOut[0])
}

func TestFanout_RangeContained(t *testing.T) {
	input := []string{
		"seeds: 79 14", //79-93
		"",
		"seed-to-soil map:",
		"52 50 48", // 50-98 map to 52-100
		"",
	}
	a := parseAlmanac(input)
	srs := getSeedRangesFromSeeds(a.seeds)
	fannedOut := a.fanoutIngredientRanges("seed", "soil", srs)
	// all of this should be matched in the 52 50 48 range.
	assert.Equal(t, 1, len(fannedOut))
	assert.Equal(t, 81, fannedOut[0].beginning)
	assert.Equal(t, 14, fannedOut[0].length)
}

func lengthSum(irs []IngredientRange) int {
	sum := 0
	for _, ir := range irs {
		sum += ir.length
	}
	return sum
}
