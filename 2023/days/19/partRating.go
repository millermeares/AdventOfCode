package nineteen

import (
	"strconv"
	"strings"
)

type PartRating struct {
	x int
	m int
	a int
	s int
}

func (pr PartRating) getMatching(r rune) int {
	if r == 'x' {
		return pr.x
	}
	if r == 'm' {
		return pr.m
	}
	if r == 'a' {
		return pr.a
	}
	if r == 's' {
		return pr.s
	}
	panic("Unexpected input " + string(r))
}

func (pr PartRating) total() int {
	return pr.x + pr.a + pr.m + pr.s
}

func (c ConditionalRule) matches(pr PartRating) bool {
	propToCompare := pr.getMatching(c.toEvaluate)
	return c.matchesProp(propToCompare)
}

func (fr FallbackRule) matches(pr PartRating) bool {
	return true
}

func parsePartRating(input string) PartRating {
	input = strings.Replace(input, "{", "", -1)
	input = strings.Replace(input, "}", "", -1)
	input = strings.Replace(input, "=", "", -1)
	input = strings.Replace(input, "x", "", -1)
	input = strings.Replace(input, "m", "", -1)
	input = strings.Replace(input, "a", "", -1)
	input = strings.Replace(input, "s", "", -1)

	// assuming that all parts are present and come in the same order every time.
	list := jsSplit(input, ",")
	x, _ := strconv.Atoi(list[0])
	m, _ := strconv.Atoi(list[1])
	a, _ := strconv.Atoi(list[2])
	s, _ := strconv.Atoi(list[3])
	return PartRating{x: x, m: m, a: a, s: s}
}

func Part1(input []string) int {
	workflows, partRatings := parseInput(input)
	sum := 0
	for _, pr := range partRatings {
		if evaluatePartRating(workflows, pr) {
			sum += pr.total()
		}
	}
	return sum
}

func parseInput(input []string) (map[string]Workflow, []PartRating) {
	workflows := map[string]Workflow{}
	var partRatings []PartRating
	parsingWorkflow := true

	for _, line := range input {
		if line == "" {
			parsingWorkflow = false
			continue
		}
		if parsingWorkflow {
			workflow := parseWorkflow(line)
			workflows[workflow.name] = workflow
		} else {
			partRatings = append(partRatings, parsePartRating(line))
		}
	}
	return workflows, partRatings
}

func (w Workflow) getResultingState(pr PartRating) string {
	for i := 0; i < len(w.rules); i++ {
		if w.rules[i].matches(pr) {
			return w.rules[i].matchedState()
		}
	}
	panic("no rules matched")
}

// returns true if reached "R" and false if reaches "A"
func evaluatePartRating(workflows map[string]Workflow, pr PartRating) bool {
	currentState := "in"
	for currentState != "A" && currentState != "R" {
		currentWorkflow := workflows[currentState]
		currentState = currentWorkflow.getResultingState(pr)
	}
	return currentState == "A"
}

// each range should have a length of 1.
func (pr PartRating) getRanges() Ranges {
	rs := map[rune]Range{}
	rs['x'] = Range{min: pr.x, max: pr.x + 1}
	rs['m'] = Range{min: pr.m, max: pr.m + 1}
	rs['a'] = Range{min: pr.a, max: pr.a + 1}
	rs['s'] = Range{min: pr.s, max: pr.s + 1}
	return Ranges{rs: rs}
}
