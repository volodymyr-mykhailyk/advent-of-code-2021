package sonar

import "github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"

func BuildVentMap(vents []ocean.Vent) ocean.VentMap {
	var ventMap = ocean.VentMap{}
	for _, vent := range vents {
		mapVent(vent, ventMap)
	}
	return ventMap
}

func mapVent(vent ocean.Vent, ventMap ocean.VentMap) {
	x1 := vent.Start.X
	x2 := vent.End.X
	y1 := vent.Start.Y
	y2 := vent.End.Y

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


