package days

import (
	"fmt"
)

func SolveConsideringOnlyNumbers(input []string) {
	solve(input, getCalibrationValue)
}

func SolveConsideringLetters(input []string) {
	solve(input, getCalibrationValueWithLetters)
}

func solve(input []string, calibrator func(input string) int) {
	total := 0
	for _, line := range input {
		calibratedValue := calibrator(line)
		total += calibratedValue
	}

	fmt.Println("Calibration value: ", total)
}
