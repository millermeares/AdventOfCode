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
	if isInvalidArrangement(input, broken) {
		// if there are more
		return 0
	}
	unknownIdx := firstUnknownIndex(input)
	if unknownIdx == -1 {
		return 1 // this is a valid combination :)
	}
	workingSpringCase := replaceAtIndex(input, '.', unknownIdx)
	workingSpringValidCount := countValidSpringArrangements(workingSpringCase, broken)
	brokenSpringCase := replaceAtIndex(input, '#', unknownIdx)
	brokenSpringValidCount := countValidSpringArrangements(brokenSpringCase, broken)
	return workingSpringValidCount + brokenSpringValidCount
}

// func updateMemo(unmatchedInputIdx int, unmatchedBrokenIdx int, val int, memo map[string]map[int]int, input string) {
// 	unmatchedInput := input[unmatchedInputIdx:]
// 	_, inputIdxSeen := memo[unmatchedInput]
// 	if !inputIdxSeen {
// 		memo[unmatchedInput] = map[int]int{}
// 	}
// 	fmt.Println("Updating memo with val", val, "at", unmatchedInput, ",", unmatchedBrokenIdx, "for input", input)
// 	memo[unmatchedInput][unmatchedBrokenIdx] = val
// }

// returns first index of unmatched for string input and broken array.
// func getUnmatched(input string, broken []int) (int, int) {
// 	unmatchedInputIdx := 0
// 	unmatchedBrokenIdx := 0
// 	// i need to know where in the input each of these strings exists. hm.
// 	brokenOrUnknownSprings := strings.Fields(strings.Replace(input, ".", " ", -1))
// 	for i := 0; i < len(brokenOrUnknownSprings); i++ {
// 		if unmatchedBrokenIdx == len(broken) || containsUnknown(brokenOrUnknownSprings[i]) {
// 			break // if there are no more brokens to match or any ? in this set, cannot match, so break.
// 		}
// 		if len(brokenOrUnknownSprings[i]) == broken[unmatchedBrokenIdx] {
// 			unmatchedBrokenIdx++
// 			unmatchedInputIdx = getIndexOfNthSpringSet(input, brokenOrUnknownSprings, i)
// 			//todo: do i have to handle when inputIdx reached len(input)? maybe not, due to current isInvalid implementation.
// 		}
// 	}
// 	return unmatchedInputIdx, unmatchedBrokenIdx
// }

// func getValue(k1 string, k2 int, m map[string]map[int]int) (int, bool) {
// 	nested, e := m[k1]
// 	if !e {
// 		return -1, false
// 	}
// 	v2, e := nested[k2]
// 	if !e {
// 		return -1, false
// 	}
// 	return v2, true
// }

// // return the index of the Nth spring set in the input.
// func getIndexOfNthSpringSet(input string, brokenSprings []string, brokenSpringIdx int) int {
// 	inputIdx := 0
// 	for i := 0; i < len(brokenSprings); i++ {
// 		indexOfMatch := indexOfSubstring(input[inputIdx:], brokenSprings[i])
// 		inputIdx = inputIdx + indexOfMatch + len(brokenSprings[i]) // inputIdx + (index of match in substring) + len(matched)
// 	}
// 	return inputIdx
// }

// func indexOfSubstring(str, subStr string) int {
// 	for i := 0; i < len(str); i++ {
// 		if str[i:i+len(subStr)] == subStr {
// 			return i
// 		}
// 	}
// 	return -1
// }

func isInvalidArrangement(input string, broken []int) bool {
	// ok so this is all that remains?
	// if it contains a ?, return false? for now.
	if containsUnknown(input) {
		return false // don't mark invalid yet -
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
