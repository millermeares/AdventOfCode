package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplestCalibration(t *testing.T) {
	actual := getCalibrationValue("42")
	assert.Equal(t, actual, 42)
}

func TestFirstDigit(t *testing.T) {
	actual := firstDigit("2")
	assert.Equal(t, int(actual), 2)
}

func TestExample1(t *testing.T) {
	actual := getCalibrationValue("1abc2")
	assert.Equal(t, actual, 12)
}

func TestExample2(t *testing.T) {
	actual := getCalibrationValue("pqr3stu8vwx")
	assert.Equal(t, actual, 38)
}

func TestExample3(t *testing.T) {
	actual := getCalibrationValue("a1b2c3d4e5f")
	assert.Equal(t, actual, 15)
}

func TestExample4(t *testing.T) {
	actual := getCalibrationValue("treb7uchet")
	assert.Equal(t, actual, 77)
}
