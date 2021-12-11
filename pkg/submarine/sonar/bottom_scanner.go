package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
)

func BuildVentMap(vents []ocean.Vent) ocean.VentMap {
	var ventMap = ocean.VentMap{}
	for _, vent := range vents {
		if vent.IsStraight() {
			mapVent(vent, ventMap)
		}
	}
	return ventMap
}

func DangerousPoints(ventMap ocean.VentMap) []structures.Point {
	var points []structures.Point
	for point, vents := range ventMap {
		if vents > 1 {
			points = append(points, point)
		}
	}
	return points
}

func mapVent(vent ocean.Vent, ventMap ocean.VentMap) {
	points := vent.Path()
	for _, point := range points {
		ventMap[point]++
	}
}
