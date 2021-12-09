package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
)

func LowPoints(hMap ocean.HeightMap) []int {
	var points []int
	hMap.AllPoints(func(height int, x int, y int) {
		if lowestPoint(height, hMap.Surroundings(x, y)) {
			points = append(points, height)
		}
	})
	return points
}

func RiskLevel(points []int) int {
	risk := 0
	for _, height := range points {
		risk += height + 1
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
