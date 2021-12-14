package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/hull"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day14")
	lines := tasks.ReadLines("cmd/day14/input.txt")
	sequence := tasks.ReadLines("cmd/day14/sequence.txt")[0]
	factory := hull.BuildPolymerFactory(lines)

	fmt.Printf("Sequencing polymer: %v\n", sequence)
	polymer := hull.SequencePolymers(sequence, factory, 10)
	fmt.Printf("Sequenced polymer info: %v\n", hull.PolymerSequenceInfo(polymer))
}
