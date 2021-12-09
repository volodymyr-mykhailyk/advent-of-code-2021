package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/sonar"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day09")
	lines := tasks.ReadLines("cmd/day09/input.txt")
	heightMap := ocean.HeightMapFromStrings(lines)

	fmt.Printf("Analyzing %vx%v height map\n", len(heightMap), len(heightMap[0]))
	riskLevel := sonar.RiskLevel(sonar.LowPoints(heightMap))

	fmt.Printf("Risk Level is: %v\n", riskLevel)
}
