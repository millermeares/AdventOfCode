package fourteen

import (
	"days"
	"fmt"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "14")
}

func Part1(input []string) int {
	tilted := tiltNorth(input)
	return calculateLoad(tilted)
}

func Part2(input []string) int {
	memo := map[string][]string{}
	for i := 0; i < 1000000000; i++ {
		if i%1000000 == 0 {
			fmt.Println("Doing cycle", i)
		}
		input = spinCycleMemo(input, memo)
	}
	printInput("after a billion spins", input)
	return calculateLoad(input)
}

// if there are a lot of duplicates, this would help.
func spinCycleMemo(input []string, memo map[string][]string) []string {
	key := inputToString(input)
	res, exists := memo[key]
	if exists {
		return res
	}
	input = spinCycle(input)
	memo[key] = input
	return input
}

func inputToString(input []string) string {
	return strings.Join(input, ",")
}

func spinCycle(input []string) []string {
	input = tiltNorth(input)
	// printInput("after north tilt", input)
	input = tiltWest(input)
	// printInput("after west tilt", input)
	input = tiltSouth(input)
	// printInput("after south tilt", input)
	input = tiltEast(input)
	// printInput("after east tilt", input)
	return input
}

func printInput(msg string, input []string) {
	fmt.Println("Printing", msg)
	for _, line := range input {
		fmt.Println(line)
	}
	fmt.Println()
}

func tiltNorth(input []string) []string {
	width := len(input[0])
	for i := 0; i < width; i++ {
		column := stringFromColumn(input, i)
		tiltedColumn := tiltRowLeft(column)
		input = replaceColumn(input, i, tiltedColumn)
	}
	return input
}

func tiltWest(input []string) []string {
	for i := 0; i < len(input); i++ {
		input[i] = tiltRowLeft(input[i])
	}
	return input
}

func tiltEast(input []string) []string {
	for i := 0; i < len(input); i++ {
		input[i] = tiltRowRight(input[i])
	}
	return input
}

func tiltSouth(input []string) []string {
	width := len(input[0])
	for i := 0; i < width; i++ {
		column := stringFromColumn(input, i)
		tiltedColumn := tiltRowRight(column)
		input = replaceColumn(input, i, tiltedColumn)
	}
	return input
}

func stringFromColumn(input []string, col int) string {
	str := ""
	for i := 0; i < len(input); i++ {
		str += string(input[i][col])
	}
	return str
}

func tiltRowLeft(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] != 'O' {
			continue
		}
		for j := i; j-1 >= 0 && input[j-1] == '.'; j-- {
			// we know input[j] is 'O' and that input[j-1] is '.' - switch them.
			input = replaceAtIndex(input, '.', j)
			input = replaceAtIndex(input, 'O', j-1)
		}
	}
	return input
}

func tiltRowRight(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != 'O' {
			continue
		}
		for j := i; j+1 < len(input) && input[j+1] == '.'; j++ {
			// we know input[j] is 'O' and that input[j+1] is '.' - switch them.
			input = replaceAtIndex(input, '.', j)
			input = replaceAtIndex(input, 'O', j+1)
		}
	}
	return input
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func replaceColumn(input []string, colIdx int, column string) []string {
	if len(input[colIdx]) != len(column) {
		panic("Unexpected column length")
	}
	for i := 0; i < len(column); i++ {
		input[i] = replaceAtIndex(input[i], rune(column[i]), colIdx)
	}
	return input
}

func calculateLoad(input []string) int {
	sum := 0
	for i, line := range input {
		weightPerRock := len(line) - i
		for _, c := range line {
			if c == 'O' {
				sum += weightPerRock
			}
		}
	}
	return sum
}
