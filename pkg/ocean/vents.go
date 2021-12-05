package ocean

import (
	"strings"
)

type Vent struct {
	Start Point
	End Point
}

type VentMap map[Point]int

func (vent Vent) IsStraight() bool {
	return vent.IsHorizontal() || vent.IsVertical()
}

func (vent Vent) IsHorizontal() bool {
	return vent.Start.X == vent.End.X
}

func (vent Vent) IsVertical() bool {
	return vent.Start.Y == vent.End.Y
}

func VentFromString(input string) Vent {
	points := strings.Split(input, " -> ")
	return Vent{Start: PointFromString(points[0]), End: PointFromString(points[1])}
}

