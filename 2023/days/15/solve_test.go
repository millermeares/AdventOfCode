package fifteen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := []string{"HASH"}
	hashed := Hash(input)
	assert.Equal(t, 52, hashed)

	input = []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}
	assert.Equal(t, 1320, Hash(input))
}

func TestPart2_Example(t *testing.T) {
	input := []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}
	power := DeriveFocusingPower(input)
	assert.Equal(t, 145, power)
}
