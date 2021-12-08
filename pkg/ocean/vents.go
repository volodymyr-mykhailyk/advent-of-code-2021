package ocean

import (
	"math"
	"strconv"
	"strings"
)

type Vent struct {
	Start Point
	End   Point
}

type VentMap map[Point]int

func (vent Vent) IsStraight() bool {
	return vent.IsHorizontal() || vent.IsVertical() || vent.IsDiagonal()
}

func (vent Vent) IsVertical() bool {
	return vent.Start.X == vent.End.X
}

func (vent Vent) IsHorizontal() bool {
	return vent.Start.Y == vent.End.Y
}

func (vent Vent) IsDiagonal() bool {
	dx := vent.Start.X - vent.End.X
	dy := vent.Start.Y - vent.End.Y
	return math.Abs(float64(dx)) == math.Abs(float64(dy))
}

func (vent Vent) XDistance() int {
	return int(math.Abs(float64(vent.Start.X - vent.End.X)))
}

func (vent Vent) YDistance() int {
	return int(math.Abs(float64(vent.Start.Y - vent.End.Y)))
}

func (vent Vent) Length() int {
	if vent.IsHorizontal() || vent.IsDiagonal() {
		return vent.XDistance()
	} else {
		return vent.YDistance()
	}
}

func (vent Vent) Path() []Point {
	length := vent.Length()
	xs := coordsRange(vent.Start.X, vent.End.X, length)
	ys := coordsRange(vent.Start.Y, vent.End.Y, length)
	points := make([]Point, length+1)
	for i := 0; i <= length; i++ {
		points[i] = Point{X: xs[i], Y: ys[i]}
	}
	return points
}

func (ventMap VentMap) Presentation() []string {
	maxX, maxY := ventMap.Dimensions()
	presentation := make([]string, maxY+1)
	for y, _ := range presentation {
		presentation[y] = strings.Repeat(".", maxX+1)
	}
	for point, count := range ventMap {
		s := presentation[point.Y]
		presentation[point.Y] = s[:point.X] + strconv.Itoa(count) + s[point.X+1:]
	}
	return presentation
}

func (ventMap VentMap) Dimensions() (int, int) {
	i := 0
	xs := make([]int, len(ventMap))
	ys := make([]int, len(ventMap))
	for point, _ := range ventMap {
		xs[i] = point.X
		ys[i] = point.Y
		i++
	}
	_, maxX := minMax(xs)
	_, maxY := minMax(ys)
	return maxX, maxY
}

func VentFromString(input string) Vent {
	points := strings.Split(input, " -> ")
	return Vent{Start: PointFromString(points[0]), End: PointFromString(points[1])}
}

func sortCoords(c1 int, c2 int) (int, int) {
	if c1 > c2 {
		return c2, c1
	} else {
		return c1, c2
	}
}

func coordsRange(s int, e int, length int) []int {
	points := make([]int, length+1)
	for i := 0; i <= length; i++ {
		if s > e {
			points[i] = s - i
		} else if s < e {
			points[i] = s + i
		} else {
			points[i] = s
		}
	}
	return points
}

func minMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
