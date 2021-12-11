package ocean

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
)

func SimulateStep(plot structures.FlatValuePlot) int {
	increaseEnergy(plot)
	flashed := make(map[structures.Point]bool)
	flash(plot, flashed)
	return len(flashed)
}

func SimulateColony(plot structures.FlatValuePlot, steps int) int {
	flashes := 0
	for i := 0; i < steps; i++ {
		step := SimulateStep(plot)
		flashes += step
	}
	return flashes
}

func SimulateGroupFlash(plot structures.FlatValuePlot) int {
	step, flashed := 0, 0
	colonySize := plot.Size()
	for flashed < colonySize {
		flashed = SimulateStep(plot)
		step++
	}
	return step
}

func flash(plot structures.FlatValuePlot, flashed map[structures.Point]bool) {
	for hasHotDumbos(plot, flashed) {
		plot.VisitAll(func(value int, point structures.Point) {
			if isReadyToFlash(value, point, flashed) {
				flashed[point] = true
				plot.VisitAround(point, func(value int, point structures.Point) {
					plot.Set(point, value+1)
				})
			}
		})
	}
	plot.VisitAll(func(value int, point structures.Point) {
		if flashed[point] {
			plot.Set(point, 0)
		}
	})
}

func hasHotDumbos(plot structures.FlatValuePlot, flashed map[structures.Point]bool) bool {
	for y, line := range plot {
		for x, value := range line {
			if isReadyToFlash(value, structures.Point{X: x, Y: y}, flashed) {
				return true
			}
		}
	}
	return false
}

func isReadyToFlash(value int, point structures.Point, flashed map[structures.Point]bool) bool {
	if value > 9 && !flashed[point] {
		return true
	} else {
		return false
	}
}

func increaseEnergy(plot structures.FlatValuePlot) {
	plot.VisitAll(func(value int, point structures.Point) {
		plot.Set(point, value+1)
	})
}
