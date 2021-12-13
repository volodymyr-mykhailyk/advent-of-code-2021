package structures

import (
	"strconv"
	"strings"
)

type AddressablePlot map[Point]int

func (plot AddressablePlot) Presentation() string {
	maxX, maxY := plot.Size()
	presentation := make([]string, maxY+1)
	for y := range presentation {
		presentation[y] = strings.Repeat(".", maxX+1)
	}
	for point, count := range plot {
		s := presentation[point.Y]
		presentation[point.Y] = s[:point.X] + strconv.Itoa(count) + s[point.X+1:]
	}
	return strings.Join(presentation, "\n")
}

func (plot AddressablePlot) Size() (int, int) {
	maxX := -1
	maxY := -1
	for point := range plot {
		if maxX < point.X {
			maxX = point.X
		}
		if maxY < point.Y {
			maxY = point.Y
		}
	}
	return maxX, maxY
}
