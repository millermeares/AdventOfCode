package main

import (
	"bufio"
	"days/thirteen"
	"fmt"
	"os"
)

func main() {
	// whenver you want to run a different day, change this line.
	day := thirteen.GetDay()

	input, err := readInputIntoList(day.DayNum())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	part1Answer := day.Part1(input)
	fmt.Println("Part 1:", part1Answer)
	part2Answer := day.Part2(input)
	fmt.Println("Part 2:", part2Answer)
}

func readInputIntoList(day string) ([]string, error) {
	file, err := os.Open("../days/" + day + "/input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}
