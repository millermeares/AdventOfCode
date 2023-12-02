package one

func toNumberSet(input string) string {
	transformed := ""
	possibilities := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(input); i++ {
		matchIdx := matchStartingAtIndex(possibilities, input, i)
		if matchIdx == -1 {
			continue
		}
		transformed = transformed + stringDigitFromPossibilities(possibilities, matchIdx)
	}
	return transformed
}

func matchStartingAtIndex(possibilities []string, total string, idx int) int {
	// iterate over possibilities
	// try to match from total[idx:end] to possibility.
	// if match, return the possibility idx
	// if not match return -1
	for possIdx, possibility := range possibilities {
		if idx+len(possibility) > len(total) {
			continue
		}
		toEval := total[idx : idx+len(possibility)]
		if possibility == toEval {
			return possIdx
		}
	}
	return -1
}

func stringDigitFromPossibilities(possibilities []string, idx int) string {
	return possibilities[digitFromPossibilities(possibilities, idx)]
}

func digitFromPossibilities(possibilities []string, idx int) int {
	if idx <= 9 {
		return idx
	}
	// this only works because of ordering of possibilities.
	return idx - 9
}

func getCalibrationValueWithLetters(input string) int {
	transformedInput := toNumberSet(input)
	return getCalibrationValue(transformedInput)
}
