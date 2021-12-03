package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/computer"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day03")
	data := tasks.ReadLines("cmd/day03/input.txt")
	fmt.Printf("Analyzing %v computer reports\n", len(data))
	fmt.Printf("Power Consumption is %v\n", computer.PowerConsumption(data))
}
