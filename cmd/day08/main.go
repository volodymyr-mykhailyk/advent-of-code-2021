package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/computer"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"strconv"
	"strings"
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
	digitsCount := 0
	simpleDigits := []string{"1", "4", "7", "8"}
	for _, line := range displayOutput {
		for _, digit := range simpleDigits {
			digitsCount += strings.Count(line, digit)
		}
	}

	fmt.Printf("Detected %v simple digits\n", digitsCount)
	sum := 0
	for _, line := range displayOutput {
		digit, _ := strconv.Atoi(line)
		sum += digit
	}

	fmt.Printf("Sum of all digits: %v\n", sum)
}
