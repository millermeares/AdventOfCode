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

func TestFanout(t *testing.T) {
	input := []string{
		"seeds: 48 52",
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
	assert.Equal(t, 46, location) // not accurate.

	// what would i expect for 48 52?
	// {48, 2}, the front not mapped to anything
	// {50, 2}, matching range 50 98 2
	// {52, 48}, matching range 50 52 48
}

func TestFanout2(t *testing.T) {
	input := []string{
		"seeds: 48 52",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
	}
	a := parseAlmanac(input)
	seedRanges := getSeedRangesFromSeeds(a.seeds)
	irs := a.fanoutIngredientRanges("seed", "soil", seedRanges)
	fmt.Println(irs)
	ir1 := irs[0]
	assert.Equal(t, 48, ir1.beginning)
	assert.Equal(t, 2, ir1.length)
}
