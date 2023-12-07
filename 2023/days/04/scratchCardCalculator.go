package four

import "days"

func GetDay() days.Day {
	return days.MakeDay(SolveScratchCardSum, SolveScratchCardAccumulator, "04")
}

func SolveScratchCardSum(input []string) int {
	sum := 0
	for _, line := range input {
		sc := parseScratchCard(line)
		sum += sc.calculatePointValue()
	}
	return sum
}
