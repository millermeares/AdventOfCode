package nineteen

import (
	"days"
	"fmt"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "19")
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

func Part2(input []string) int {
	return -1
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

// returns true if reached "R" and false if reaches "A"
func evaluatePartRating(workflows map[string]Workflow, pr PartRating) bool {
	currentState := "in"
	for currentState != "A" && currentState != "R" {
		currentWorkflow := workflows[currentState]
		currentState = currentWorkflow.getResultingState(pr)
		fmt.Println("At state", currentState, "after evaluating workflow", currentWorkflow.name)
	}
	return currentState == "A"
}

type Workflow struct {
	name  string
	rules []Rule
}

func (w Workflow) getResultingState(pr PartRating) string {
	for i := 0; i < len(w.rules); i++ {
		if w.rules[i].matches(pr) {
			return w.rules[i].matchedState()
		}
	}
	panic("no rules matched")
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

type Rule interface {
	matches(pr PartRating) bool
	matchedState() string
}

type ConditionalRule struct {
	operator    rune
	compareTo   int
	stateIfTrue string
	toEvaluate  rune
}

func (c ConditionalRule) matches(pr PartRating) bool {
	propToCompare := pr.getMatching(c.toEvaluate)
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

func (fr FallbackRule) matches(pr PartRating) bool {
	return true
}

func (fr FallbackRule) matchedState() string {
	return fr.state
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

func jsSplit(input string, substring string) []string {
	input = strings.Replace(input, substring, " ", -1)
	return strings.Fields(input)
}
