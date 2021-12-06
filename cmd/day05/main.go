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
	var straightVents []ocean.Vent
	tasks.Iterate(lines, func(input string, i int) {
		vent := ocean.VentFromString(input)
		vents = append(vents, vent)
		if vent.IsStraight() {
			straightVents = append(straightVents, vent)
		}
	})

	fmt.Printf("Analyzing %v vents\n", len(vents))
	fmt.Printf("Mapping %v straight vents\n", len(straightVents))

	straightVentsMap := sonar.BuildVentMap(straightVents)
	straightDangerousPoints := sonar.DangerousPoints(straightVentsMap)

	fmt.Printf("Dangerous Points %v", len(straightDangerousPoints))
}

