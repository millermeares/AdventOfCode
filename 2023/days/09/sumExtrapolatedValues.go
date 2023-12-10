package nine

import (
	"days"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(SolveExtrapolatedValues, Part2, "09")
}

func SolveExtrapolatedValues(input []string) int {
	sets := parseInput(input)
	sum := 0
	for _, set := range sets {
		sum += calculateNextValue(set)
	}
	return sum
}

func calculateNextValue(set []int) int {
	if allZeros(set) {
		return 0
	}
	var diffSet []int
	for i := 0; i < len(set)-1; i++ {
		diff := set[i+1] - set[i]
		diffSet = append(diffSet, diff)
	}
	return set[len(set)-1] + calculateNextValue(diffSet)
}

func allZeros(set []int) bool {
	for _, num := range set {
		if num != 0 {
			return false
		}
	}
	return true
}

func reverseArray(numbers []int) []int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func Part2(input []string) int {
	sets := parseInput(input)
	sum := 0
	for _, set := range sets {
		reversed := reverseArray(set)
		extrapolated := calculateNextValue(reversed)
		sum += extrapolated
	}
	return sum
}

func parseInput(input []string) [][]int {
	var sets [][]int
	for _, line := range input {
		var set []int
		for _, v := range strings.Fields(line) {
			parsed, _ := strconv.Atoi(v)
			set = append(set, parsed)
		}
		sets = append(sets, set)
	}
	return sets
}
