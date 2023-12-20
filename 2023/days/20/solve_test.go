package twenty

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	input := []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}
	fmt.Println(input)
	assert.Equal(t, 32000000, Part1(input))
}

func TestExample1_2(t *testing.T) {
	input := []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}
	fmt.Println(input)
	assert.Equal(t, 11687500, Part1(input))
}

func TestPart2(t *testing.T) {

}
