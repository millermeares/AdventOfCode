package fifteen

import (
	"days"
	"strconv"
	"strings"
)

func GetDay() days.Day {
	return days.MakeDay(Hash, DeriveFocusingPower, "15")
}

func Hash(input []string) int {
	input = commaDelimitedIntoList(input)
	sum := 0
	for _, line := range input {
		sum += getHash(line)
	}
	return sum
}

func commaDelimitedIntoList(input []string) []string {
	return strings.Fields(strings.Replace(input[0], ",", " ", -1))
}

func getHash(input string) int {
	val := 0
	for _, c := range input {
		code := int(c)
		val += code
		val *= 17
		val = val % 256
	}
	return val
}

func DeriveFocusingPower(input []string) int {
	input = commaDelimitedIntoList(input)
	// there are 256 boxes through which light will pass.
	// along wall, there is a lens library with focal lengths from 1 to 9
	// letters at beginning of each step indicate the label of the lens on which the step operates
	// label followed by "-" or "=", which indicate which operation to perform.
	// if -, go to box and remove any lens with the given label in the box (if it exists). move the other lenses in the box "forward" as far as they will go.
	// if =, use the following number as the focal length.
	// 		if there is already a lens with the same label, replace the old lens with the new lens and don't move any lenses in the box
	// 		if there is not already a lens in the box with the label, add the lens behind all other lenses in the box (or at the front if needed).

	boxes := initBoxes()
	for _, op := range input {
		label := labelFromOperation(op)
		box := boxes[getHash(label)]
		if strings.Contains(op, "-") {
			box.removeLens(label)
		} else {
			box.replaceLens(label, op)
		}
	}
	return sumBoxPower(boxes)
}

func sumBoxPower(boxes []*Box) int {
	sum := 0
	for i, box := range boxes {
		sum += box.calcluateFocusingPower(i)
	}
	return sum
}

func labelFromOperation(input string) string {
	return strings.Fields(strings.Replace(strings.Replace(input, "-", " ", -1), "=", " ", -1))[0]
}

type Box struct {
	lenses []Lens
}

func (b *Box) indexOfLens(label string) int {
	for i := 0; i < len(b.lenses); i++ {
		if b.lenses[i].label == label {
			return i
		}
	}
	return -1
}

func (b *Box) removeLens(label string) {
	lensToRemove := b.indexOfLens(label)
	if lensToRemove == -1 {
		return // does not exist in box, do nothing.
	}
	b.lenses = append(b.lenses[:lensToRemove], b.lenses[lensToRemove+1:]...)
}

func (b *Box) replaceLens(label string, op string) {
	focalLength, e := strconv.Atoi(strings.Split(op, "=")[1])
	if e != nil {
		panic(e)
	}
	newLens := Lens{label: label, strength: focalLength}
	existingLensIndex := b.indexOfLens(label)
	if existingLensIndex != -1 {
		b.lenses[existingLensIndex] = newLens
	} else {
		b.lenses = append(b.lenses, newLens)
	}
}

func (b Box) calcluateFocusingPower(index int) int {
	multiplier := index + 1
	sum := 0
	for i, lens := range b.lenses {
		sum += (multiplier * (i + 1) * lens.strength)
	}
	return sum
}

type Lens struct {
	label    string
	strength int
}

func initBoxes() []*Box {
	var boxes []*Box
	for i := 0; i < 256; i++ {
		boxes = append(boxes, &Box{
			lenses: make([]Lens, 0),
		})
	}
	return boxes
}
