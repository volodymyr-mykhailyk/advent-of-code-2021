package science

func SumElements(elements []int) int {
	total := 0
	for _, value := range elements {
		total += value
	}
	return total
}
func MultiplyElements(elements []int) int {
	total := 1
	for _, value := range elements {
		total *= value
	}
	return total
}
