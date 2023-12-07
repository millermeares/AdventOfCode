package days

type Day interface {
	Part1(input []string) int
	Part2(input []string) int
	DayNum() string // functionally, folder name.
}

type GenericDay struct {
	part1 func(input []string) int
	part2 func(input []string) int
	num   string
}

func (d GenericDay) Part1(input []string) int {
	return d.part1(input)
}

func (d GenericDay) Part2(input []string) int {
	return d.part2(input)
}

func (d GenericDay) DayNum() string {
	return d.num
}

func MakeDay(part1 func(input []string) int, part2 func(input []string) int, num string) GenericDay {
	return GenericDay{
		part1: part1,
		part2: part2,
		num:   num,
	}
}
