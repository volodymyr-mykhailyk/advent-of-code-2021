package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"sort"
)

func LowPoints(hMap ocean.HeightMap) []ocean.Point {
	var points []ocean.Point
	hMap.VisitAll(func(height int, point ocean.Point) {
		if lowestPoint(height, hMap.Surroundings(point)) {
			points = append(points, point)
		}
	})
	return points
}

func LargestBasins(hMap ocean.HeightMap, points []ocean.Point) [][]ocean.Point {
	basins := make([][]ocean.Point, len(points))
	for i, point := range points {
		basins[i] = basinFor(hMap, point)
	}
	return basins
}

func LowPointsRiskLevel(hMap ocean.HeightMap, points []ocean.Point) int {
	risk := 0
	for _, point := range points {
		risk += hMap.Get(point) + 1
	}
	return risk
}

func BasinsRiskLevel(hMap ocean.HeightMap, basins [][]ocean.Point) int {
	sizes := make([]int, len(basins))
	for i, basin := range basins {
		sizes[i] = len(basin)
	}
	sort.IntSlice(sizes).Sort()
	largestBasins := sizes[len(sizes)-3:]
	risk := 1
	for _, size := range largestBasins {
		risk *= size
	}
	return risk
}

func lowestPoint(height int, surroundings []int) bool {
	for _, h := range surroundings {
		if h <= height {
			return false
		}
	}
	return true
}

func basinFor(hMap ocean.HeightMap, point ocean.Point) []ocean.Point {
	var basinPoints = make(map[ocean.Point]bool)
	basinAround(hMap, point, basinPoints)
	points := make([]ocean.Point, len(basinPoints))
	i := 0
	for point, _ := range basinPoints {
		points[i] = point
		i++
	}
	return points
}

func basinAround(hMap ocean.HeightMap, point ocean.Point, visited map[ocean.Point]bool) {
	visited[point] = true
	hMap.VisitAround(point, func(height int, p ocean.Point) {
		if !visited[p] && height < 9 {
			basinAround(hMap, p, visited)
		}
	})
}
