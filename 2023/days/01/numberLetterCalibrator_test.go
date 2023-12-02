package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibrationWithLetters_Example1(t *testing.T) {
	actual := getCalibrationValueWithLetters("two1nine")
	assert.Equal(t, 29, actual)
}

func TestCalibrationWithLetters_Example2(t *testing.T) {
	actual := getCalibrationValueWithLetters("nine1two8")
	assert.Equal(t, 98, actual)
}

func TestCalibration_SingleWord(t *testing.T) {
	actual := getCalibrationValueWithLetters("two")
	assert.Equal(t, 22, actual)
}

func TestCalibrationWithLetters_Example3(t *testing.T) {
	runCalibrationTest(t, "abcone2threexyz", 13)
}

func TestCalibrationWithLetters_Example4(t *testing.T) {
	runCalibrationTest(t, "xtwone3four", 24)
}

func runCalibrationTest(t *testing.T, input string, expected int) {
	actual := getCalibrationValueWithLetters(input)
	assert.Equal(t, expected, actual)
}
