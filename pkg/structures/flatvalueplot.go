package structures

import (
	"math"
	"strconv"
	"strings"
)

type FlatValuePlot [][]int

func (plot FlatValuePlot) Get(point Point) int {
	if plot.IsInbound(point) {
		return plot[point.Y][point.X]
	} else {
		return math.MaxInt
	}
}

func (plot FlatValuePlot) Set(point Point, value int) {
	if plot.IsInbound(point) {
		plot[point.Y][point.X] = value
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

func (plot FlatValuePlot) VisitAll(f func(value int, point Point)) {
	for y, row := range plot {
		for x, value := range row {
			f(value, Point{X: x, Y: y})
		}
	}
}

func (plot FlatValuePlot) VisitAdjacent(point Point, f func(value int, point Point)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.Left(), f)
}

func (plot FlatValuePlot) VisitAround(point Point, f func(value int, point Point)) {
	plot.VisitPoint(point.Up(), f)
	plot.VisitPoint(point.UpRight(), f)
	plot.VisitPoint(point.Right(), f)
	plot.VisitPoint(point.DownRight(), f)
	plot.VisitPoint(point.Down(), f)
	plot.VisitPoint(point.DownLeft(), f)
	plot.VisitPoint(point.Left(), f)
	plot.VisitPoint(point.UpLeft(), f)
}

func (plot FlatValuePlot) VisitPoint(point Point, f func(value int, point Point)) {
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

func (plot FlatValuePlot) Printable(valueSep string) []string {
	result := make([]string, len(plot))
	for y, line := range plot {
		elements := make([]string, len(line))
		for x, value := range line {
			elements[x] = strconv.FormatInt(int64(value), 10)
		}
		result[y] = strings.Join(elements, valueSep)
	}
	return result
}

func (plot FlatValuePlot) Size() int {
	return len(plot) * len(plot[0])
}

func FlatValuesFromString(input [][]string) FlatValuePlot {
	maxY := len(input)
	maxX := len(input[0])
	plot := make(FlatValuePlot, maxY)
	for y, line := range input {
		plot[y] = make([]int, maxX)
		for x, value := range line {
			h, _ := strconv.ParseInt(value, 10, 0)
			plot[y][x] = int(h)
		}
	}
	return plot
}
