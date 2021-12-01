package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/science"
)

func LinearIncreaseSpeed(readings []int) int {
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

func AveragedIncreaseSpeed(readings []int) int {
	var averagedDepth []int
	for i := 0; i < (len(readings) - 2); i++ {
		averagedDepth = append(averagedDepth, science.SumElements(readings[i:i+3]))
	}
	return LinearIncreaseSpeed(averagedDepth)
}
