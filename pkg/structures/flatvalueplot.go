package structures

import (
	"math"
	"strconv"
)

type FlatValuePlot [][]int

func (plot FlatValuePlot) Get(point Point) int {
	if plot.IsInbound(point) {
		return plot[point.Y][point.X]
	} else {
		return math.MaxInt
	}
}

func (plot FlatValuePlot) Surroundings(point Point) []int {
	results := make([]int, 4)
	results[0] = plot.Get(point.Up())
	results[1] = plot.Get(point.Right())
	results[2] = plot.Get(point.Down())
	results[3] = plot.Get(point.Left())
	return results
}

func (plot FlatValuePlot) VisitAll(f func(height int, point Point)) {
	for y, row := range plot {
		for x, height := range row {
			f(height, Point{X: x, Y: y})
		}
	}
}

func (plot FlatValuePlot) VisitAround(point Point, f func(height int, point Point)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.Left(), f)
}

func (plot FlatValuePlot) VisitPoint(point Point, f func(height int, point Point)) {
	if plot.IsInbound(point) {
		f(plot.Get(point), point)
	}
}

func (plot FlatValuePlot) IsInbound(point Point) bool {
	x, y := point.X, point.Y
	if y < 0 || y >= len(plot) {
		return false
	} else if x < 0 || x >= len(plot[y]) {
		return false
	} else {
		return true
	}
}

func FlatValuesFromString(input []string) FlatValuePlot {
	maxY := len(input)
	maxX := len(input[0])
	plot := make(FlatValuePlot, maxY)
	for y, line := range input {
		plot[y] = make([]int, maxX)
		for x, height := range line {
			h, _ := strconv.ParseInt(string(height), 10, 0)
			plot[y][x] = int(h)
		}
	}
	return plot
}
