package twelve

import (
	"strconv"
	"strings"
)

func countSpringArrangements(input string) int {
	split := strings.Split(input, " ")
	broken := getIntList(strings.Split(split[1], ","))
	return countValidSpringArrangements(split[0], broken)
}

func countValidSpringArrangements(input string, broken []int) int {
	input = moveToFirstNonDot(input)

	// base case.
	if !containsUnknown(input) {
		return oneIfValid(input, broken)
	}
	if canPrune(input, broken) {
		return 0
	}
	dotIdx := firstDotIndex(input)
	unknownIdx := firstUnknownIndex(input)

	if dotIdx != -1 && dotIdx < unknownIdx {
		// if there is a dot before unknown index, we can short circuit the problem.
		matchedThrough := getMatchedThrough(input[:dotIdx], broken)
		if matchedThrough == -1 {
			// could be included in the pruning function.
			return 0
		}
		return countValidSpringArrangements(input[dotIdx:], broken[matchedThrough:])
	}

	brokenSpringCase := replaceAtIndex(input, '#', unknownIdx)

	springBrokenCount := 0
	// we currently match the next set of broken things.
	if len(broken) > 0 && broken[0] == unknownIdx+1 {
		// fmt.Println("New # at index", unknownIdx, "in input", brokenSpringCase, "matched set", brokenSpringCase[:unknownIdx+1], "which should be of length", broken[0])
		// input[0:unknownIdx+1] is "#", *AND* it matches broken[0].
		// now there are 4 cases.
		if unknownIdx == len(input)-1 {
			// 1. we are at the end of the string. springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
			springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
		} else {
			nextChar := input[unknownIdx+1]
			if nextChar == '.' {
				// 2. input[unknownIdx+1] is a period. springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
				springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
			} else if nextChar == '#' {
				// 3. if input[unknownIdx+1] is a #, not a valid match no further evaluation needed.
				springBrokenCount = 0
			} else if nextChar == '?' {
				// 4. if input[unknownIdx+1] is a ?, then unknownIdx+1 *cannot* be a #. Set it to a period and then call countValidSpringArrangements(replacedBrokenSpring, broken[1:])
				brokenSpringCase = replaceAtIndex(brokenSpringCase, '.', unknownIdx+1)
				springBrokenCount = countValidSpringArrangements(brokenSpringCase[unknownIdx+1:], broken[1:])
				// springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
			}
		}
	} else {
		springBrokenCount = countValidSpringArrangements(brokenSpringCase, broken)
	}

	// move to first non-# and then try to match a "broken"? to reduce search space?
	// careful of ##?? 2 1  because ###. is not valid. but if you greedily take ##, then ?? would match "11" twice, rather than once.
	// need to explicitly handle if "?" is coming next.

	// creating two separate problems. in theory, memoization would be useful for this part.
	workingSpringCase := replaceAtIndex(input, '.', unknownIdx)
	matchedThrough := getMatchedThrough(workingSpringCase[:unknownIdx], broken)
	workingSpringCount := 0
	if matchedThrough != -1 {
		// only enter this block if the matching process was valid.
		workingSpringCount = countValidSpringArrangements(workingSpringCase[unknownIdx:], broken[matchedThrough:])
	}

	return springBrokenCount + workingSpringCount
}

func canPrune(input string, broken []int) bool {
	// input isn't long enough to fulfill broken's "needs"
	requiredBroken, minDots := getBrokenRequirements(broken)
	if len(input) < requiredBroken+minDots {
		return true
	}
	// too many broken provisioned.
	if countBroken(input) > requiredBroken {
		return true
	}
	return false
}

func getBrokenRequirements(broken []int) (int, int) {
	// specific amount of # required
	requiredBroken := sum(broken)
	minDots := len(broken) - 1
	return requiredBroken, minDots
}

func sum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func moveToFirstNonDot(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] != '.' {
			return input[i:]
		}
	}
	return ""
}

func countBroken(input string) int {
	return strings.Count(input, "#")
}

// returns -1 if it is invalid and 0 should be returned.
func getMatchedThrough(input string, broken []int) int {
	if containsUnknown(input) {
		panic("Not expecting any unknown in this matching function")
	}
	// can len(broken) ever be smaller than len(brokenSprings)?
	matchedThrough := 0
	brokenSprings := strings.Fields(strings.Replace(input, ".", " ", -1))

	if len(broken) < len(brokenSprings) {
		return -1 // does not match.
	}
	for i, springSet := range brokenSprings {
		if broken[i] != len(springSet) {
			// -1 indicates that it is not a match.
			return -1
		}
		matchedThrough++
	}
	return matchedThrough
}

func oneIfValid(input string, broken []int) int {
	matchedThrough := getMatchedThrough(input, broken)
	if matchedThrough == len(broken) {
		return 1
	}
	return 0
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

func firstDotIndex(input string) int {
	for i, c := range input {
		if c == '.' {
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
