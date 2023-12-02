package days

import (
	"bufio"
	"fmt"
	"os"
)

func SolveConsideringOnlyNumbers() {
	solve(getCalibrationValue)
}

func SolveConsideringLetters() {
	solve(getCalibrationValueWithLetters)
}

func solve(calibrator func(input string) int) {
	fmt.Println("Solving calibration")
	input, err := readInputIntoList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	total := 0

	for _, line := range input {
		calibratedValue := calibrator(line)
		total += calibratedValue
	}

	fmt.Println("Calibration value: ", total)
}

func readInputIntoList() ([]string, error) {
	file, err := os.Open("../days/01/input.txt")
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
