package navigation

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"math"
)

type Path []structures.Point

func (path Path) End() structures.Point {
	return path[len(path)-1]
}

func (path Path) ContinueWith(point structures.Point) Path {
	return append(path[0:len(path):len(path)], point)
}

func (path Path) HasPoint(point structures.Point) bool {
	for _, p := range path {
		if p == point {
			return true
		}
	}
	return false
}

func (path Path) LengthAt(plot structures.FlatValuePlot) int {
	length := 0
	for i, point := range path {
		if i > 0 {
			length += plot.Get(point)
		}
	}
	return length
}

func BuildRiskLevelPlot(plot structures.FlatValuePlot, start structures.Point, end structures.Point) structures.FlatValuePlot {
	maxX, maxY := plot.Dimensions()
	distancePlot := structures.FlatValuesFromSpec(maxX, maxY, math.MaxInt)
	buildDistancePlotV2(plot, distancePlot, start, end)
	return distancePlot
}

func ScaleRiskPlot(plot structures.FlatValuePlot, size int) structures.FlatValuePlot {
	maxX, maxY := plot.Dimensions()
	newPlot := structures.FlatValuesFromSpec(maxX*size, maxY*size, 0)
	plot.VisitAll(func(value int, point structures.Point) {
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				newValue := value + x + y
				for newValue > 9 {
					newValue -= 9
				}
				newPoint := structures.Point{X: point.X + (maxX * x), Y: point.Y + (maxY * y)}
				//fmt.Printf("Point %v, %v\n", point, newPoint)
				newPlot.Set(newPoint, newValue)
			}
		}
	})
	return newPlot
}

func buildDistancePlot(plot structures.FlatValuePlot, distancePlot structures.FlatValuePlot, path Path, end structures.Point) {
	currentPoint := path.End()
	currentRisk := path.LengthAt(plot)
	previousRisk := distancePlot.Get(currentPoint)
	if previousRisk > currentRisk {
		distancePlot.Set(currentPoint, currentRisk)
		if currentPoint == end {
			return
		} else {
			plot.VisitAdjacent(currentPoint, func(value int, point structures.Point) {
				if !path.HasPoint(point) {
					buildDistancePlot(plot, distancePlot, path.ContinueWith(point), end)
				}
			})
		}
	}
}

func buildDistancePlotV2(plot structures.FlatValuePlot, distancePlot structures.FlatValuePlot, start, end structures.Point) {
	visitQueue := make([]Path, 0)
	visitQueue = append(visitQueue, Path{start})
	i := 0
	for len(visitQueue) > 0 {
		i++
		currentPath := visitQueue[0]
		visitQueue = visitQueue[1:]
		currentPoint := currentPath.End()
		currentRisk := currentPath.LengthAt(plot)
		previousRisk := distancePlot.Get(currentPoint)
		//fmt.Printf("Visiting point %v, %v, %v\n", currentPoint, currentRisk, currentPath)
		if previousRisk > currentRisk {
			//fmt.Printf("Better path at %v, %v, %v\n", currentPoint, currentRisk, currentPath)
			distancePlot.Set(currentPoint, currentRisk)
			if currentPoint != end {
				plot.VisitAdjacent(currentPoint, func(value int, point structures.Point) {
					if !currentPath.HasPoint(point) {
						//fmt.Printf("Planing visit: %v, %v\n", point, currentPath.ContinueWith(point))
						visitQueue = append(visitQueue, currentPath.ContinueWith(point))
						//fmt.Printf("Next Queue: %v\n", visitQueue)
					}
				})
			} else {
				//fmt.Printf("Reached end iteration: %v, risk: %v, path: %v\n", i, currentRisk, currentPath)
			}
		}
	}
	fmt.Printf("Total Iterations %v\n", i)
}

func LowestRiskPathLevel(plot structures.FlatValuePlot, end structures.Point) int {
	return plot.Get(end)
}
