package vision

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"strconv"
	"strings"
)

func VisibleActivationDots(plot structures.AddressablePlot) int {
	return len(plot)
}

func FoldActivationSheet(sheet structures.AddressablePlot, instruction string) structures.AddressablePlot {
	fold := getFoldFunction(instruction)
	return fold(sheet)
}

func getFoldFunction(instruction string) func(sheet structures.AddressablePlot) structures.AddressablePlot {
	parts := strings.Split(instruction, "=")
	at, _ := strconv.Atoi(parts[1])
	function := foldX
	if strings.Contains(instruction, "fold along y=") {
		function = foldY
	}
	return func(sheet structures.AddressablePlot) structures.AddressablePlot {
		return function(sheet, at)
	}
}

func foldY(sheet structures.AddressablePlot, at int) structures.AddressablePlot {
	newSheet := make(structures.AddressablePlot, len(sheet))
	for point := range sheet {
		x, y := point.Coords()
		if point.Y > at {
			y = 2*at - y
		}
		newSheet[structures.Point{X: x, Y: y}] = 1
	}
	return newSheet
}

func foldX(sheet structures.AddressablePlot, at int) structures.AddressablePlot {
	newSheet := make(structures.AddressablePlot, len(sheet))
	for point := range sheet {
		x, y := point.Coords()
		if point.X > at {
			x = 2*at - x
		}
		newSheet[structures.Point{X: x, Y: y}] = 1
	}
	return newSheet
}

func ActivationSheet(input []string) structures.AddressablePlot {
	sheet := make(structures.AddressablePlot, len(input))
	for _, line := range input {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		sheet[structures.Point{X: x, Y: y}] = 1
	}
	return sheet
}
