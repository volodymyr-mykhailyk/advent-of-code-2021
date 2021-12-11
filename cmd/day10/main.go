package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/navigation"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day10")
	lines := tasks.ReadLines("cmd/day10/input.txt")

	fmt.Printf("Analyzing %v of navigation lines\n", len(lines))
	syntaxErrorScore := navigation.SystemErrorScore(lines)

	fmt.Printf("System Error Score: %v\n", syntaxErrorScore)

	autoCompleteScore := navigation.SystemAutocompleteScore(lines)
	fmt.Printf("System Auto Complete Score: %v\n", autoCompleteScore)
}
