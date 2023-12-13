package twelve

import (
	"strconv"
	"strings"
)

func countSpringArrangements(input string) int {
	split := strings.Split(input, " ")
	broken := getIntList(strings.Split(split[1], ","))
	memo := map[State]int{}
	return countValidSpringArrangements(split[0], broken, memo)
}

func countValidSpringArrangements(input string, broken []int, memo map[State]int) int {
	if !containsUnknown(input) {
		return oneIfValid(input, broken)
	}
	// state := calculateState(input, broken)
	// m, e := memo[state]
	// if e {
	// 	return m
	// }

	children := getChildInputs(input)
	sum := 0
	for _, child := range children {
		// if i was going to prune, i would do it here. i am not sure that is the best way.
		sum += countValidSpringArrangements(child, broken, memo)
	}

	memo[calculateState(input, broken)] = sum //todo: if i am going to memo, it is going to be here.
	return sum
}
func getChildInputs(input string) []string {
	unknownIdx := firstUnknownIndex(input)
	workingSpringCase := replaceAtIndex(input, '.', unknownIdx)
	brokenSpringCase := replaceAtIndex(input, '#', unknownIdx)
	return []string{workingSpringCase, brokenSpringCase}
}

func oneIfValid(input string, broken []int) int {
	brokenSprings := strings.Fields(strings.Replace(input, ".", " ", -1))
	if len(brokenSprings) != len(broken) {
		return 0
	}
	for i, springSet := range brokenSprings {
		if len(springSet) != broken[i] {
			return 0
		}
	}
	return 1
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func firstUnknownIndex(input string) int {
	for i, c := range input {
		if c == '?' {
			return i
		}
	}
	return -1
}

func containsUnknown(input string) bool {
	return strings.Contains(input, "?")
}

func getIntList(input []string) []int {
	var nums []int
	for _, num := range input {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}
	return nums
}

type State struct {
}

// maybe it's by prefix? like "matchedSoFar"? does that make sense? no. because all prefixes are theoretically unique.
func calculateState(input string, broken []int) State {
	return State{}
}
