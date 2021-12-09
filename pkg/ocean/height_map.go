package ocean

import (
	"math"
	"strconv"
)

type HeightMap [][]int

func (hMap HeightMap) Get(x int, y int) int {
	if y < 0 || y >= len(hMap) {
		return math.MaxInt
	} else if x < 0 || x >= len(hMap[y]) {
		return math.MaxInt
	} else {
		return hMap[y][x]
	}
}

func (hMap HeightMap) Surroundings(x int, y int) []int {
	results := make([]int, 4)
	results[0] = hMap.Get(x, y-1)
	results[1] = hMap.Get(x+1, y)
	results[2] = hMap.Get(x, y+1)
	results[3] = hMap.Get(x-1, y)
	return results
}

func (hMap HeightMap) AllPoints(f func(height int, x int, y int)) {
	for y, row := range hMap {
		for x, height := range row {
			f(height, x, y)
		}
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
