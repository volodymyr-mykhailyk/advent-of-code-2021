package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/sonar"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
)

func main() {
	tasks.Announce("Day01")
	data := transformations.ParseIntegers(tasks.ReadLines("cmd/day01/input.txt"))
	fmt.Printf("Scanning %v sonar readings\n", len(data))
	fmt.Printf("Depth increasing with %v linear speed\n", sonar.LinearIncreaseSpeed(data))
	fmt.Printf("Depth increasing with %v averaged speed\n", sonar.AveragedIncreaseSpeed(data))
}
