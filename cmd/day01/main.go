package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day01")
	data := tasks.ParseIntegers(tasks.ReadLines("cmd/day01/input.txt"))
	fmt.Printf("Scanning %v sonar readings\n", len(data))
	depthIncreases := submarine.FloorIncreaseSpeed(data)
	fmt.Printf("Depth increasing with %v speed\n", depthIncreases)
}
