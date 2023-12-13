package main

import (
	"bufio"
	"days/twelve"
	"fmt"
	"os"
	"time"
)

func main() {
	// whenver you want to run a different day, change this line.
	day := twelve.GetDay()

	input, err := readInputIntoList(day.DayNum())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	timeFunc("Part 1", input, day.Part1)
	timeFunc("Part 2", input, day.Part2)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s time: %v\n", name, time.Since(start))
	}
}

func timeFunc(name string, input []string, f func([]string) int) {
	defer timer(name)()
	part1Answer := f(input)
	fmt.Println(name, "answer:", part1Answer)
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
