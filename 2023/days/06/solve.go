package six

import "days"

func GetDay() days.Day {
	return days.MakeDay(SolveWaysToBeatRecord, SolveWaysToBeatRecordSingleRace, "06")
}

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
