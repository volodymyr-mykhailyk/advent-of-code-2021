package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/vision"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day13")
	lines := tasks.ReadLines("cmd/day13/input.txt")
	instructions := tasks.ReadLines("cmd/day13/instructions.txt")
	sheet := vision.ActivationSheet(lines)

	fmt.Printf("Rading activation sheet with %v points\n", vision.VisibleActivationDots(sheet))
	sheet = vision.FoldActivationSheet(sheet, instructions[0])
	fmt.Printf("Dots after first fold: %v\n", vision.VisibleActivationDots(sheet))
}
