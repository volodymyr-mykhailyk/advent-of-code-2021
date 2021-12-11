package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
)

func main() {
	tasks.Announce("Day11")
	lines := tasks.ReadLines("cmd/day11/input.txt")
	colony := structures.FlatValuesFromString(transformations.SplitLines(lines, ""))

	fmt.Printf("Simulating %vx%v octopus colony\n", len(colony), len(colony[0]))
	flashes := ocean.SimulateColony(colony, 100)
	fmt.Printf("Total number of dumbos flashed: %v\n", flashes)
}
