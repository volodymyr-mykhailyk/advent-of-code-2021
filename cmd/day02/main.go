package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/navigation"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day02")
	commands := tasks.ReadLines("cmd/day02/input.txt")
	fmt.Printf("Predicting %v submatrine movements\n", len(commands))
	coordinates := navigation.PredictPosition(commands, submarine.Coordinates{})
	distanceChecksum := coordinates.Position * coordinates.Depth
	fmt.Printf("Predicted position: %v. Calculated distance checksum: %v\n", coordinates, distanceChecksum)
}
