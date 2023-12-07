package two

import (
	"days"
	"strconv"
	"strings"
)

const (
	Red   = "red"
	Blue  = "blue"
	Green = "green"
)

func GetDay() days.Day {
	return days.MakeDay(Part1, Part2, "02")
}

func addColorToMap(colorCounts map[string]int, rawColor string) {
	colorCount := strings.Split(rawColor, " ")
	parsedCount, _ := strconv.Atoi(colorCount[0]) // swallow error on purpose
	colorCounts[colorCount[1]] = parsedCount
}

func Part1(input []string) int {
	games := parseGames(input)
	return solvePossibleGames(games)
}

func solvePossibleGames(games []Game) int {
	qualifyingGames := getQualifyingGames(games)
	idSum := getSumOfGameIds(qualifyingGames)
	return idSum
}

// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
// prediction: these values will be input for part 2.
func getQualifyingGames(games []Game) []Game {
	var qualifyingGames []Game
	for _, game := range games {
		maxRed := game.maxValueOfColor(Red)
		maxGreen := game.maxValueOfColor(Green)
		maxBlue := game.maxValueOfColor(Blue)
		if maxRed > 12 || maxGreen > 13 || maxBlue > 14 {
			continue
		}
		qualifyingGames = append(qualifyingGames, game)
	}
	return qualifyingGames
}

func getSumOfGameIds(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.id
	}
	return sum
}
