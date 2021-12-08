package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/sonar"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day05")
	lines := tasks.ReadLines("cmd/day05/input.txt")
	var vents []ocean.Vent
	var perpendicularVents []ocean.Vent
	tasks.Iterate(lines, func(input string, i int) {
		vent := ocean.VentFromString(input)
		vents = append(vents, vent)
		if !vent.IsDiagonal() {
			perpendicularVents = append(perpendicularVents, vent)
		}
	})

	fmt.Printf("Analyzing %v vents\n", len(vents))
	fmt.Printf("Mapping %v perpendicular vents\n", len(perpendicularVents))

	perpendicularVentsMap := sonar.BuildVentMap(perpendicularVents)
	straightDangerousPoints := sonar.DangerousPoints(perpendicularVentsMap)

	fmt.Printf("Dangerous Points: %v\n", len(straightDangerousPoints))

	fmt.Printf("Mapping all %v vents\n", len(vents))

	ventsMap := sonar.BuildVentMap(vents)
	dangerousPoints := sonar.DangerousPoints(ventsMap)

	fmt.Printf("Dangerous Points: %v", len(dangerousPoints))
}

