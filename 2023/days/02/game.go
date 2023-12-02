package two

import (
	"strconv"
	"strings"
)

type Game struct {
	id       int
	handfuls []Handful
}

func buildGame(input string) Game {
	parsing := strings.Split(input, ":")
	id, _ := strconv.Atoi(strings.Split(parsing[0], " ")[1])
	handfuls := buildHandfuls(parsing[1])
	return Game{
		id:       id,
		handfuls: handfuls,
	}
}

func (g Game) maxValueOfColor(color string) int {
	max := 0
	for _, handful := range g.handfuls {
		val := handful.getValueOfColor(color)
		if val > max {
			max = val
		}
	}
	return max
}

func parseGames(input []string) []Game {
	var games []Game
	for _, line := range input {
		games = append(games, buildGame(line))
	}
	return games
}

func (g Game) fewestCubeCountPower() int {
	reqGreen := g.maxValueOfColor(Green)
	reqRed := g.maxValueOfColor(Red)
	reqBlue := g.maxValueOfColor(Blue)
	return reqGreen * reqBlue * reqRed
}
