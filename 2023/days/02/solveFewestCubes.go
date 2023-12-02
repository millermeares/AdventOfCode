package two

import "fmt"

func SolveFewestCubes(input []string) {
	games := parseGames(input)
	power := solveFewestCubesPower(games)
	fmt.Println("Fewest cubes power: ", power)
}

func solveFewestCubesPower(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.fewestCubeCountPower()
	}
	return sum
}
