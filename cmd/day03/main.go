package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/computer"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
)

func main() {
	tasks.Announce("Day03")
	data := transformations.ParseBits(tasks.ReadLines("cmd/day03/input.txt"))
	fmt.Printf("Analyzing %v computer reports\n", len(data))
	fmt.Printf("Power Consumption is %v\n", computer.PowerConsumption(data, 12))
	fmt.Printf("Life Support is %v\n", computer.LifeSupportRating(data, 12))
}
