package ocean

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (point Point) Up() Point {
	return Point{X: point.X, Y: point.Y - 1}
}

func (point Point) Right() Point {
	return Point{X: point.X + 1, Y: point.Y}
}

func (point Point) Down() Point {
	return Point{X: point.X, Y: point.Y + 1}
}

func (point Point) Left() Point {
	return Point{X: point.X - 1, Y: point.Y}
}

func PointFromString(input string) Point {
	coordinates := transformations.ParseIntegers(strings.Split(input, ","))
	return Point{X: coordinates[0], Y: coordinates[1]}
}
