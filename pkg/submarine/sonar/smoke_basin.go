package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"sort"
)

func LowPoints(plot structures.FlatValuePlot) []structures.Point {
	var points []structures.Point
	plot.VisitAll(func(height int, point structures.Point) {
		if lowestPoint(height, plot.Surroundings(point)) {
			points = append(points, point)
		}
	})
	return points
}

func LargestBasins(plot structures.FlatValuePlot, points []structures.Point) [][]structures.Point {
	basins := make([][]structures.Point, len(points))
	for i, point := range points {
		basins[i] = basinFor(plot, point)
	}
	return basins
}

func LowPointsRiskLevel(plot structures.FlatValuePlot, points []structures.Point) int {
	risk := 0
	for _, point := range points {
		risk += plot.Get(point) + 1
	}
	return risk
}

func BasinsRiskLevel(_ structures.FlatValuePlot, basins [][]structures.Point) int {
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

func basinFor(plot structures.FlatValuePlot, point structures.Point) []structures.Point {
	var basinPoints = make(map[structures.Point]bool)
	basinAround(plot, point, basinPoints)
	points := make([]structures.Point, len(basinPoints))
	i := 0
	for point := range basinPoints {
		points[i] = point
		i++
	}
	return points
}

func basinAround(plot structures.FlatValuePlot, point structures.Point, visited map[structures.Point]bool) {
	visited[point] = true
	plot.VisitAround(point, func(height int, p structures.Point) {
		if !visited[p] && height < 9 {
			basinAround(plot, p, visited)
		}
	})
}
