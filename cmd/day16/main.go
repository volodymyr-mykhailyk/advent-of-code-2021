package main

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine/computer"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
)

func main() {
	tasks.Announce("Day15")
	input := tasks.ReadLines("cmd/day16/input.txt")[0]
	fmt.Printf("Reading packet of size: %v\n", len(input))

	packet := computer.ReadBitsPacket(input)
	fmt.Printf("Version Checksum: %v\n", computer.BitsPacketVersionChecksum(packet))
	fmt.Printf("Value: %v\n", packet.Value())
}
