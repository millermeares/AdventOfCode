package nineteen

import (
	"days"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "19")
}

func parameterizedPart2(input []string, beginRange Ranges) int {
	workflows, _ := parseInput(input)
	count := 0
	queue := []RangeDestination{
		{destination: "in", ranges: beginRange},
	}
	for len(queue) > 0 {
		pq, popped := dequeue(queue)
		queue = pq
		if popped.destination == "R" {
			continue // doesn't match any.
		}
		if popped.destination == "A" {
			count += popped.ranges.getCount()
			continue
		}

		workflow := workflows[popped.destination]
		ranges := popped.ranges
		for _, rule := range workflow.rules {
			matchedRange, unmatchedRange := rule.splitRanges(ranges)
			if matchedRange.getCount() > 0 {
				queue = append(queue, RangeDestination{
					destination: rule.matchedState(), ranges: matchedRange,
				})
			}
			if unmatchedRange.getCount() == 0 {
				break // no unmatched, so no need to continue rule evaluation
			}
			ranges = unmatchedRange // set to new value to let next rule be evaluated.
		}
	}
	return count
}

func Part2(input []string) int {
	return parameterizedPart2(input, starterRanges(4000))
}

func dequeue(rds []RangeDestination) ([]RangeDestination, RangeDestination) {
	popped := rds[len(rds)-1]
	return rds[:len(rds)-1], popped
}

type RangeDestination struct {
	ranges      Ranges
	destination string
}

type Workflow struct {
	name  string
	rules []Rule
}

func parseWorkflow(input string) Workflow {
	idSplit := jsSplit(input, "{")
	id := idSplit[0]
	rawRules := strings.Replace(idSplit[1], "}", "", -1)
	rules := parseRules(rawRules)
	return Workflow{name: id, rules: rules}
}

func parseRules(input string) []Rule {
	var rules []Rule
	rawRules := jsSplit(input, ",")
	for _, rule := range rawRules {
		rules = append(rules, parseRule(rule))
	}
	return rules
}

func parseRule(input string) Rule {
	isConditional := strings.Contains(input, ":")
	if isConditional {
		return parseConditionalRule(input)
	}
	// fallback should just be the string of the final state.
	return FallbackRule{state: input}
}

func parseConditionalRule(input string) ConditionalRule {
	propToEvaluate := input[0]
	operator := input[1]
	compareToResultingState := jsSplit(input[2:], ":")
	resultingState := compareToResultingState[1]
	compareTo, _ := strconv.Atoi(compareToResultingState[0])
	return ConditionalRule{
		toEvaluate:  rune(propToEvaluate),
		operator:    rune(operator),
		stateIfTrue: resultingState,
		compareTo:   compareTo,
	}
}

type Rule interface {
	matches(pr PartRating) bool
	matchesProp(propToCompare int) bool
	matchedState() string
	splitRanges(r Ranges) (Ranges, Ranges)
}

type ConditionalRule struct {
	toEvaluate  rune
	operator    rune
	compareTo   int
	stateIfTrue string
}

func (c ConditionalRule) matchesProp(propToCompare int) bool {
	if c.operator == '>' {
		return propToCompare > c.compareTo
	}
	if c.operator == '<' {
		return propToCompare < c.compareTo
	}
	panic("Unexpected operator: " + string(c.operator))
}

func (c ConditionalRule) matchedState() string {
	return c.stateIfTrue
}

type FallbackRule struct {
	state string
}

func (fr FallbackRule) matchedState() string {
	return fr.state
}

func (fr FallbackRule) matchesProp(propToCompare int) bool {
	return true
}

func (fr FallbackRule) splitRanges(rs Ranges) (Ranges, Ranges) {
	return rs, emptyRanges()
}

func jsSplit(input string, substring string) []string {
	input = strings.Replace(input, substring, " ", -1)
	return strings.Fields(input)
}

type Ranges struct {
	rs map[rune]Range
}

type Range struct {
	min int
	max int
}

func (r Range) length() int {
	return r.max - r.min
}

// returns matchingRange, nonMatchingRange
func (cr ConditionalRule) splitRanges(ranges Ranges) (Ranges, Ranges) {
	propRange := ranges.rs[cr.toEvaluate]
	matchesMin := cr.matchesProp(propRange.min)
	matchesMax := cr.matchesProp(propRange.max - 1) // subtract because actual max is excluded.
	// if matches both, the entire range is matched.
	if matchesMax && matchesMin {
		return ranges, emptyRanges() // rule includes this range
	}
	if !matchesMax && !matchesMin {
		return emptyRanges(), ranges // rule does not include this range.
	}

	if matchesMin {
		matchedRange := Range{min: propRange.min, max: cr.compareTo}
		notMatchedRange := Range{min: cr.compareTo, max: propRange.max}
		return Ranges{rs: copyWithValue(ranges.rs, cr.toEvaluate, matchedRange)},
			Ranges{rs: copyWithValue(ranges.rs, cr.toEvaluate, notMatchedRange)}
	}
	matchedRange := Range{min: cr.compareTo + 1, max: propRange.max}
	notMatchedRange := Range{min: propRange.min, max: cr.compareTo + 1}
	return Ranges{rs: copyWithValue(ranges.rs, cr.toEvaluate, matchedRange)},
		Ranges{rs: copyWithValue(ranges.rs, cr.toEvaluate, notMatchedRange)}
}

func (r Ranges) getCount() int {
	product := 1
	for _, r := range r.rs {
		product *= r.length()
	}
	return product
}

func starterRanges(maximum int) Ranges {
	ranges := map[rune]Range{}
	ranges['m'] = Range{min: 1, max: maximum + 1}
	ranges['s'] = Range{min: 1, max: maximum + 1}
	ranges['a'] = Range{min: 1, max: maximum + 1}
	ranges['x'] = Range{min: 1, max: maximum + 1}
	return Ranges{rs: ranges}
}

func emptyRanges() Ranges {
	ranges := map[rune]Range{}
	ranges['m'] = Range{min: 0, max: 0}
	ranges['s'] = Range{min: 0, max: 0}
	ranges['a'] = Range{min: 0, max: 0}
	ranges['x'] = Range{min: 0, max: 0}
	return Ranges{rs: ranges}
}

func copy(ranges map[rune]Range) map[rune]Range {
	copy := map[rune]Range{}
	for k, v := range ranges {
		copy[k] = v
	}
	return copy
}

func copyWithValue(ranges map[rune]Range, k rune, v Range) map[rune]Range {
	copy := copy(ranges)
	copy[k] = v
	return copy
}
