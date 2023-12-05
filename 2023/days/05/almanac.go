package five

import (
	"strconv"
	"strings"
)

// numbers and seeds are used to identify fertizilizer, water, and seed type
// seeds that need to be planted: seeds.

// each line is a range.
// 50, 98, 2
// 50=>98, 51=>99
type Almanac struct {
	seeds     []int
	rangeMaps []*RangeMap
}

type RangeMap struct {
	to     string
	from   string
	ranges []Range
}

type Range struct {
	destination int
	source      int
	length      int
	raw         string
}

func (a Almanac) getMapping(source string, destination string, sourceValue int) int {
	if source == destination {
		return sourceValue // source is destination, done.
	}
	// todo: figure out return case. maybe when source = destination?
	// this relies on the idea that there is only one per source.
	rm := a.getRangeMapForSource(source)
	destinationValue := rm.getDestinationMapping(sourceValue)
	return a.getMapping(rm.to, destination, destinationValue)
}

func (a *Almanac) getRangeMapForSource(source string) *RangeMap {
	for _, rm := range a.rangeMaps { // maybe a.rangeMaps is being copied here?
		if rm.from == source {
			return rm
		}
	}
	panic("No range map for source: " + source)
}

// they want the lowest location number that corresponds
func parseAlmanac(input []string) Almanac {
	a := Almanac{}
	var recentMapSource string
	for _, line := range input {
		if strings.Contains(line, "seeds:") {
			a.seeds = parseSeeds(line)
			continue
		}
		if strings.Contains(line, "map:") {
			rm := initRangeMap(line)
			recentMapSource = rm.from
			a.addRangeMap(rm) // ok the problem is that
			continue
		}
		if len(line) == 0 {
			continue
		}
		r := parseRange(line)
		rm := a.getRangeMapForSource(recentMapSource)
		rm.ranges = append(rm.ranges, r)
	}
	return a
}

func parseSeeds(input string) []int {
	rawSeeds := strings.Fields(strings.Split(input, ":")[1])
	var seeds []int
	for _, rawSeed := range rawSeeds {
		seed, _ := strconv.Atoi(rawSeed)
		seeds = append(seeds, seed)
	}
	return seeds
}

func (a *Almanac) addRangeMap(rm RangeMap) {
	a.rangeMaps = append(a.rangeMaps, &rm)
}

func initRangeMap(input string) RangeMap {
	mapping := strings.Split(strings.Split(input, " ")[0], "-")
	return RangeMap{
		to:     mapping[2],
		from:   mapping[0],
		ranges: []Range{},
	}
}

func (rm *RangeMap) addRange(r Range) {
	rm.ranges = append(rm.ranges, r)
}

func parseRange(input string) Range {
	fields := strings.Fields(input)
	destination, _ := strconv.Atoi(fields[0])
	source, _ := strconv.Atoi(fields[1])
	length, _ := strconv.Atoi(fields[2])
	return Range{
		destination: destination,
		source:      source,
		length:      length,
		raw:         input,
	}
}

func (rm RangeMap) getDestinationMapping(val int) int {
	for _, r := range rm.ranges {
		if !r.inMappingRange(val) {
			continue
		}
		return r.getMapping(val)
	}
	return val // if none in range, default to same value.
}

func (r Range) inMappingRange(sourceVal int) bool {
	minSource := r.source
	maxSource := r.source + r.length
	return sourceVal >= minSource && sourceVal < maxSource
}

func (r Range) getMapping(sourceVal int) int {
	// only call if already in mapping range.
	diff := sourceVal - r.source
	return r.destination + diff
}
