package six

import (
	"strconv"
	"strings"
)

type Race struct {
	time           int
	recordDistance int
}

func parseInput(input []string) []Race {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distances := strings.Fields(strings.Split(input[1], ":")[1])
	var races []Race
	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			time:           parseInt(times[i]),
			recordDistance: parseInt(distances[i]),
		})
	}
	return races
}

func parseInputToSingleRace(input []string) Race {
	times := strings.Split(input[0], ":")[1]
	singleTime := strings.Replace(times, " ", "", -1)
	distances := strings.Split(input[1], ":")[1]
	singleDistance := strings.Replace(distances, " ", "", -1)
	return Race{
		time:           parseInt(singleTime),
		recordDistance: parseInt(singleDistance),
	}
}

// simply swallows exception
func parseInt(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

func (r Race) amountOfWaysToBeat() int {
	// iterate over the amount of seconds that the button could be held.
	// check if would have beaten.
	waysToWin := 0
	for buttonHeldTime := 1; buttonHeldTime < r.time; buttonHeldTime++ {
		// buttonHeldTime = speed
		timeRacing := r.time - buttonHeldTime
		speed := buttonHeldTime
		distanceTraveled := timeRacing * speed
		if distanceTraveled > r.recordDistance {
			waysToWin++
		}
	}
	return waysToWin
}
