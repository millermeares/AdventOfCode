package days

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func SolveConsideringOnlyNumbers() {
	solve(getCalibrationValue)
}

func SolveConsideringLetters() {
	solve(getCalibrationValueWithLetters)
}

func solve(calibrator func(input string) int) {
	fmt.Println("Solving calibration")
	input, err := readInputIntoList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	total := 0

	for _, line := range input {
		calibratedValue := calibrator(line)
		total += calibratedValue
	}

	fmt.Println("Calibration value: ", total)
}

func getCalibrationValue(input string) int {
	base := int(firstDigit(input)) * 10
	ones := int(lastDigit(input))
	return base + ones
}

// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
func firstDigit(input string) rune {
	for _, char := range input {
		if !isDigit(char) {
			continue
		}
		return char - '0'
	}
	panic("Expected to find an integer")
}

func lastDigit(input string) rune {
	for i := len(input) - 1; i >= 0; i-- {
		char := rune(input[i])
		if !isDigit(char) {
			continue
		}
		return char - '0'
	}
	panic("Expected to find an integer")
}

func isDigit(input rune) bool {
	return unicode.IsDigit(input)
}

func readInputIntoList() ([]string, error) {
	file, err := os.Open("../days/01/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
