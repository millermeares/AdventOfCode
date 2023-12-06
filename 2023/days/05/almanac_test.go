package five

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeMap(t *testing.T) {
	rm := initRangeMap("seed-to-soil map:")
	input := []string{
		"50 98 2",
		"52 50 48",
	}
	for _, line := range input {
		r := parseRange(line)
		rm.addRange(r)
	}
	assert.Equal(t, "seed", rm.from)
	assert.Equal(t, "soil", rm.to)

	val := rm.getDestinationMapping(1)
	assert.Equal(t, 1, val)

	val = rm.getDestinationMapping(50)
	assert.Equal(t, 52, val)

	val = rm.getDestinationMapping(51)
	assert.Equal(t, 53, val)

	val = rm.getDestinationMapping(52)
	assert.Equal(t, 54, val)
}

func TestRange(t *testing.T) {
	r := parseRange("98 50 2")
	assert.Equal(t, 50, r.source)
	assert.Equal(t, 98, r.destination)
	assert.Equal(t, 2, r.length)

	assert.True(t, r.inMappingRange(50))
	assert.False(t, r.inMappingRange(52))

	val := r.getMapping(50)
	assert.Equal(t, 98, val)

	val = r.getMapping(51)
	assert.Equal(t, 99, val)

	val = r.getMapping(52)
	assert.Equal(t, 100, val)
}

func TestRange_Example(t *testing.T) {
	r := parseRange("52 50 48")
	assert.True(t, r.inMappingRange(79))
	val := r.getMapping(79)
	assert.Equal(t, 81, val)

	rm := initRangeMap("seed-to-soil map:")
	rm.addRange(r)
	val = rm.getDestinationMapping(79)
	assert.Equal(t, 81, val)
}

func TestGetMapping_Examples(t *testing.T) {
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
	a := parseAlmanac(input)
	assert.Equal(t, 7, len(a.rangeMaps))

	seed1 := a.getMapping("seed", "soil", 79)
	assert.Equal(t, 81, seed1)
}
