package two

func Part2(input []string) int {
	games := parseGames(input)
	return solveFewestCubesPower(games)
}

func solveFewestCubesPower(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.fewestCubeCountPower()
	}
	return sum
}
