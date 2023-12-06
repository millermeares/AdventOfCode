package six

func SolveWaysToBeatRecord(input []string) int {
	races := parseInput(input)
	answer := 1
	for _, race := range races {
		waysToWin := race.amountOfWaysToBeat()
		answer *= waysToWin
	}
	return answer
}

func SolveWaysToBeatRecordSingleRace(input []string) int {
	race := parseInputToSingleRace(input)
	return race.amountOfWaysToBeat()
}
