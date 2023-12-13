package twelve

import (
	"days"
	"fmt"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(SumCountSpringArrangements, SumCountFoldedSpringArrangements, "12")
}

func SumCountSpringArrangements(input []string) int {
	sum := 0
	for _, line := range input {
		count := countSpringArrangements(line)
		// fmt.Println("Calculated", count, "for line", line)
		sum += count
	}
	return sum
}

func SumCountFoldedSpringArrangements(input []string) int {
	sum := 0
	for i, line := range input {
		fmt.Println("Starting line", i)
		sum += countExpandedSpringArrangement(line)
	}
	return sum
}

func countExpandedSpringArrangement(line string) int {
	expanded := expandSpringInput(line)
	count := countSpringArrangements(expanded)
	fmt.Println("Calculated", count, "for line", line, "that was expanded to", expanded)
	return count
}

func expandSpringInput(input string) string {
	split := strings.Fields(input)
	expandedSprings := ""
	expandedBroken := ""
	for i := 0; i < 5; i++ {
		if i != 0 {
			expandedSprings += "?"
			expandedBroken += ","
		}
		expandedSprings += split[0]
		expandedBroken += split[1]
	}
	return expandedSprings + " " + expandedBroken
}