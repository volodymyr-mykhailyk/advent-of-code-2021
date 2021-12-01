package submarine

func FloorIncreaseSpeed(readings []int) int {
	increase := 0
	previous := 9223372036854775807
	for _, s := range readings {
		if previous < s {
			increase += 1
		}
		previous = s
	}
	return increase
}
