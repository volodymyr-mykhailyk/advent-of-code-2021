package ocean

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"strings"
)

type Point struct {
	X int
	Y int
}

func PointFromString(input string) Point {
	coordinates := transformations.ParseIntegers(strings.Split(input, ","))
	return Point{X: coordinates[0], Y: coordinates[1]}
}
