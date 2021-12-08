package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/computer"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day08")
	lines := tasks.ReadLines("cmd/day08/input.txt")
	entries := make([]computer.Entry, len(lines))
	tasks.Iterate(lines, func(input string, i int) {
		entries[i] = computer.ParseEntry(input)
	})


	fmt.Printf("Analyzing %v display entries\n", len(entries))

	displayOutput := computer.DisplayOutput(entries)
	simpleDigits := 0
	for _, line := range(displayOutput) {
		for _, digit := range line {
			if digit != -1 {
				simpleDigits++
			}
		}
	}

	fmt.Printf("Detected %v simple digits\n", simpleDigits)
}

