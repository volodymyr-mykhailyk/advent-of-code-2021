package sonar

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
)

func BuildVentMap(vents []ocean.Vent) ocean.VentMap {
	var ventMap = ocean.VentMap{}
	for _, vent := range vents {
		fmt.Printf("vent: %v, straight: %v\n", vent, vent.IsStraight())
		if vent.IsStraight() {
			mapVent(vent, ventMap)
		}
	}
	return ventMap
}

func DangerousPoints(ventMap ocean.VentMap) []ocean.Point {
	var points []ocean.Point
	for point, vents := range ventMap {
		if vents > 1 {
			points = append(points, point)
		}
	}
	return points
}

func mapVent(vent ocean.Vent, ventMap ocean.VentMap) {
	x1, x2 := sortCoords(vent.Start.X, vent.End.X)
	y1, y2 := sortCoords(vent.Start.Y, vent.End.Y)

	if x1 != x2 {
		for i := x1; i <= x2; i++ {
			ventMap[ocean.Point{X: i, Y: y1}]++
		}
	}
	if y1 != y2 {
		for i := y1; i <= y2; i++ {
			ventMap[ocean.Point{X: x1, Y: i}]++
		}
	}
}

func sortCoords(c1 int, c2 int) (int, int) {
	if c1 > c2 {
		return c2, c1
	} else {
		return c1, c2
	}
}


