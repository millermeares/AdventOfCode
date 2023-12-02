package one

import "unicode"

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
