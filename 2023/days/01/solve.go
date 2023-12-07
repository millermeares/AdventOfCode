package one

import "days"

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "01")
}
func Part1(input []string) int {
	return solve(input, getCalibrationValue)
}

func Part2(input []string) int {
	return solve(input, getCalibrationValueWithLetters)
}

func solve(input []string, calibrator func(input string) int) int {
	total := 0
	for _, line := range input {
		calibratedValue := calibrator(line)
		total += calibratedValue
	}
	return total
}
