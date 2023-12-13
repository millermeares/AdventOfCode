package thirteen

import (
	"days"
)

func GetDay() days.Day {
	return days.MakeDay(DetectMirrors, DetectMirrorsWithSmudge, "13")
}

func DetectMirrors(input []string) int {
	sum := 0
	sections := parseInput(input)
	for _, section := range sections {
		score, _, _ := getSectionMirrorScore(section, -1, -1)
		sum += score
	}
	return sum
}

func DetectMirrorsWithSmudge(input []string) int {
	sum := 0
	sections := parseInput(input)
	for _, section := range sections {
		sum += getSmudgedMirrorScore(section)
	}
	return sum
}

// brute force is definitely not the most efficient way to do this, but it does work.
func getSmudgedMirrorScore(section []string) int {
	vanillaScore, matchedVert, matchedHoriz := getSectionMirrorScore(section, -1, -1)
	maxForSection := 0
	for i := 0; i < len(section); i++ {
		for j := 0; j < len(section[i]); j++ {
			oldRow := section[i]
			currentCharacter := rune(section[i][j])
			newRow := replaceAtIndex(oldRow, oppositeCharacter(currentCharacter), j)
			section[i] = newRow
			iterationScore, _, _ := getSectionMirrorScore(section, matchedVert, matchedHoriz)
			section[i] = oldRow
			// fixing the smudge must give a *different* line.
			if vanillaScore != iterationScore && iterationScore > maxForSection {
				maxForSection = iterationScore
			}
		}
	}
	return maxForSection
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func oppositeCharacter(c rune) rune {
	if c == '#' {
		return '.'
	}
	return '#'
}

func getSectionMirrorScore(rows []string, dontMatchVertical int, dontMatchHorizontal int) (int, int, int) {
	verticalIndex := getIndexOfVerticalReflectionLine(rows, dontMatchVertical)
	horizontalIndex := getIndexOfVerticalReflectionLine(invert(rows), dontMatchHorizontal)
	if verticalIndex > horizontalIndex {
		return verticalIndex + 1, verticalIndex, horizontalIndex
	}
	return (horizontalIndex + 1) * 100, verticalIndex, horizontalIndex
}

// dontMatchRow is a super gross solution for part 2 that should technically work.
func getIndexOfVerticalReflectionLine(input []string, dontMatchRow int) int {
	max := -1
	width := len(input[0])
	// minus one because last cannot be the left-most point of a mirror.
	for i := 0; i < width-1; i++ {
		for row := 0; row < len(input); row++ {
			if !isMirrorPalindrome(input[row], i) {
				break // not valid
			}
			// if we made it through the last row, we should update max.
			if row == len(input)-1 && i != dontMatchRow {
				max = i
			}
		}
	}
	return max
}

func invert(input []string) []string {
	// each column should become a row.
	var inverted []string
	for i := 0; i < len(input[0]); i++ {
		inverted = append(inverted, getColumn(input, i))
	}
	return inverted
}

func getColumn(input []string, idx int) string {
	var column = ""
	for _, line := range input {
		column = column + string(line[idx])
	}
	return column
}

// idx represents left starting.
func isMirrorPalindrome(input string, idx int) bool {
	for i := 0; idx-i >= 0 && idx+i+1 < len(input); i++ {
		leftComp := idx - i
		rightComp := idx + 1 + i
		if input[leftComp] != input[rightComp] {
			return false
		}
	}
	return true
}

func parseInput(input []string) [][]string {
	var sections [][]string
	var currentRows []string
	for _, line := range input {
		if line == "" {
			sections = append(sections, currentRows)
			currentRows = []string{}
			continue
		}
		currentRows = append(currentRows, line)
	}
	sections = append(sections, currentRows)
	return sections
}
