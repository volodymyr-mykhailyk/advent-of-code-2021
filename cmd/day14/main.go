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

	fmt.Printf("Sequencing short polymer: %v\n", sequence)
	shortPolymer := hull.SequencePolymers(sequence, factory, 10)
	fmt.Printf("Sequenced short polymer info: %v\n", shortPolymer.SequenceInfo())

	fmt.Printf("Sequencing long polymer: %v\n", sequence)
	longPolymer := hull.SequencePolymers(sequence, factory, 40)
	fmt.Printf("Sequenced long polymer info: %v\n", longPolymer.SequenceInfo())
}
