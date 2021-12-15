package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/navigation"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
)

func main() {
	tasks.Announce("Day15")
	lines := tasks.ReadLines("cmd/day15/input.txt")
	plot := structures.FlatValuesFromString(transformations.SplitLines(lines, ""))

	fmt.Printf("Finding Lowest Risk Path at map of size: %v\n", plot.Size())
	riskPlot := navigation.BuildRiskLevelPlot(plot, plot.Start(), plot.End())
	riskLevel := navigation.LowestRiskPathLevel(riskPlot, plot.End())
	fmt.Printf("Safest path risk level: %v\n", riskLevel)
}
