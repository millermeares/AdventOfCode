package four

func SolveScratchCardSum(input []string) int {
	sum := 0
	for _, line := range input {
		sc := parseScratchCard(line)
		sum += sc.calculatePointValue()
	}
	return sum
}
