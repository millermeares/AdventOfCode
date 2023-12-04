package three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveEngineSchematic_Example(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 4361, sum)
}

func TestSolveEngineSchematic_OneLine(t *testing.T) {
	input := []string{"467..114.."}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 0, sum)
}

func TestSolveEngineSchematic_OneLine_Char(t *testing.T) {
	input := []string{"467.*.114.."}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 0, sum)
}

func TestSolveEngineSchematic_Right(t *testing.T) {
	input := []string{"467*..."}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 467, sum)
}

func TestSolveEngineSchematic_Left(t *testing.T) {
	input := []string{"...*467"}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 467, sum)
}

func TestSolveEngineSchematic_CharBelow(t *testing.T) {
	input := []string{
		"467..114..",
		"*....114..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 467, sum)
}

func TestSolveEngineSchema_Diagonal_TopLeft(t *testing.T) {
	input := []string{
		"467.*.....",
		".....114..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 114, sum)
}

func TestSolveEngineSchema_Diagonal_TopRight(t *testing.T) {
	input := []string{
		"467.....*.",
		".....114..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 114, sum)
}

func TestSolveEngineSchema_Diagonal_BottomLeft(t *testing.T) {
	input := []string{
		".467.......",
		"*.....114..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 467, sum)
}

func TestSolveEngineSchema_Diagonal_BottomRight(t *testing.T) {
	input := []string{
		".467.......",
		"....*.114..",
	}
	sum := SolveEngineSchematic(input)
	assert.Equal(t, 467, sum)
}
