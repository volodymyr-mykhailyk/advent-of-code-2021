package ocean

import (
	"math"
	"strconv"
)

type HeightMap [][]int

func (hMap HeightMap) Get(point Point) int {
	if hMap.IsInbound(point) {
		return hMap[point.Y][point.X]
	} else {
		return math.MaxInt
	}
}

func (hMap HeightMap) Surroundings(point Point) []int {
	results := make([]int, 4)
	results[0] = hMap.Get(point.Up())
	results[1] = hMap.Get(point.Right())
	results[2] = hMap.Get(point.Down())
	results[3] = hMap.Get(point.Left())
	return results
}

func (hMap HeightMap) VisitAll(f func(height int, point Point)) {
	for y, row := range hMap {
		for x, height := range row {
			f(height, Point{X: x, Y: y})
		}
	}
}

func (hMap HeightMap) VisitAround(point Point, f func(height int, point Point)) {
	hMap.VisitPoint(point.Up(), f)
	hMap.VisitPoint(point.Right(), f)
	hMap.VisitPoint(point.Down(), f)
	hMap.VisitPoint(point.Left(), f)
}

func (hMap HeightMap) VisitPoint(point Point, f func(height int, point Point)) {
	if hMap.IsInbound(point) {
		f(hMap.Get(point), point)
	}
}

func (hMap HeightMap) IsInbound(point Point) bool {
	x, y := point.X, point.Y
	if y < 0 || y >= len(hMap) {
		return false
	} else if x < 0 || x >= len(hMap[y]) {
		return false
	} else {
		return true
	}
}

func HeightMapFromStrings(input []string) HeightMap {
	maxY := len(input)
	maxX := len(input[0])
	hMap := make(HeightMap, maxY)
	for y, line := range input {
		hMap[y] = make([]int, maxX)
		for x, height := range line {
			h, _ := strconv.ParseInt(string(height), 10, 0)
			hMap[y][x] = int(h)
		}
	}
	return hMap
}
