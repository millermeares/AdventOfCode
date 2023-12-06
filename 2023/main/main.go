package main

import (
	"bufio"
	"days/six"
	"fmt"
	"os"
)

func main() {
	input, err := readInputIntoList("06")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	answer := six.SolveWaysToBeatRecordSingleRace(input)
	fmt.Println("Answer: ", answer)
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
