package twelve

import (
	"fmt"
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
	state := calculateState(input, broken)
	fmt.Println("Calculating state", state, "from input", input, "and broken", broken)
	m, exists := memo[state]
	if exists {
		fmt.Println("Early returning", m, "from state", state, "from memo for input", input)
		return m
	}
	if isInvalidArrangement(input, broken) {
		// if there are more
		return 0
	}
	unknownIdx := firstUnknownIndex(input)
	if unknownIdx == -1 {
		return 1 // this is a valid combination since there are no unknowns and it's not invalid.
	}
	workingSpringCase := replaceAtIndex(input, '.', unknownIdx)
	workingSpringValidCount := countValidSpringArrangements(workingSpringCase, broken, memo)

	brokenSpringCase := replaceAtIndex(input, '#', unknownIdx)
	brokenSpringValidCount := countValidSpringArrangements(brokenSpringCase, broken, memo)

	total := workingSpringValidCount + brokenSpringValidCount
	fmt.Println("Setting", total, "from input", input, "and state", state)
	memo[state] = total
	return total
}

// consider iterating through the different things to find if it's illegal?
func isInvalidArrangement(input string, broken []int) bool {
	requiredBroken := sum(broken)
	if countNonWorking(input) < requiredBroken {
		return true
	}
	if containsUnknown(input) {
		return false // only evaluate if no ?
	}
	brokenSprings := strings.Fields(strings.Replace(input, ".", " ", -1))
	if len(brokenSprings) != len(broken) {
		return true
	}
	for i, springSet := range brokenSprings {
		if len(springSet) != broken[i] {
			return true
		}
	}
	return false
}

func countNonWorking(input string) int {
	count := 0
	for _, v := range input {
		if v != '.' {
			count++
		}
	}
	return count
}

func sum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
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

// this represents the un-evaluated
type State struct {
}
