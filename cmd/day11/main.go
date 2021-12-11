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
	colony1 := structures.FlatValuesFromString(transformations.SplitLines(lines, ""))

	fmt.Printf("Simulating %v octopus colony1\n", colony1.Size())
	flashes := ocean.SimulateColony(colony1, 100)
	fmt.Printf("Total number of dumbos flashed: %v\n", flashes)

	groupFlashAt := ocean.SimulateGroupFlash(colony1) + 100
	fmt.Printf("Everyone will flash at: %v\n", groupFlashAt)
}
