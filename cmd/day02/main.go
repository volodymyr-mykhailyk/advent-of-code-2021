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
	linearCoords := navigation.PredictLinearPosition(commands, submarine.ZeroCoordinates())
	linearChecksum := linearCoords.Position * linearCoords.Depth
	fmt.Printf("Predicted Linear Position: %v. Calculated distance checksum: %v\n", linearCoords, linearChecksum)
	aimedCoords := navigation.PredictAimedPosition(commands, submarine.ZeroCoordinates())
	aimedChecksum := aimedCoords.Position * aimedCoords.Depth
	fmt.Printf("Predicted Aimed Position: %v. Calculated distance checksum: %v\n", aimedCoords, aimedChecksum)
}
