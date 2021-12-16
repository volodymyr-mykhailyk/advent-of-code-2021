package computer

import (
	"fmt"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/science"
	"sort"
	"strconv"
	"strings"
)

type BitsPacket struct {
	Version    string
	Type       string
	Contents   string
	Source     string
	SubPackets []BitsPacket
}

func (packet BitsPacket) Value() int {
	if packet.Type == "100" {
		return bitsToInt(packet.Contents)
	} else {
		values := make([]int, len(packet.SubPackets))
		for i, p := range packet.SubPackets {
			values[i] = p.Value()
		}
		return valuesFunctions[packet.Type](values)
	}
}

var valuesFunctions = map[string]func(values []int) int{
	"000": func(values []int) int {
		return science.SumElements(values)
	},
	"001": func(values []int) int {
		return science.MultiplyElements(values)
	},
	"010": func(values []int) int {
		sort.Ints(values)
		return values[0]
	},
	"011": func(values []int) int {
		sort.Ints(values)
		return values[len(values)-1]
	},
	"101": func(values []int) int {
		if values[0] > values[1] {
			return 1
		} else {
			return 0
		}
	},
	"110": func(values []int) int {
		if values[0] < values[1] {
			return 1
		} else {
			return 0
		}
	},
	"111": func(values []int) int {
		if values[0] == values[1] {
			return 1
		} else {
			return 0
		}
	},
}

func ReadBitsPacket(input string) BitsPacket {
	input = convertToBitsString(input)
	packet := BitsPacket{}
	readPacket(&packet, input)
	return packet
}

func BitsPacketVersionChecksum(packet BitsPacket) int {
	sum := bitsToInt(packet.Version)
	for _, subPacket := range packet.SubPackets {
		sum += BitsPacketVersionChecksum(subPacket)
	}
	return sum
}

func readPacket(packet *BitsPacket, input string) string {
	input = readPacketInfo(packet, input)
	input = readPacketContents(packet, input)
	return input
}

func readPacketInfo(packet *BitsPacket, input string) string {
	packet.Version, input = readNextBits(input, 3)
	packet.Type, input = readNextBits(input, 3)
	return input
}

func readPacketContents(packet *BitsPacket, input string) string {
	if packet.Type == "100" {
		return readLiteralContents(packet, input)
	} else {
		return readOperatorContents(packet, input)
	}
}

func readOperatorContents(packet *BitsPacket, input string) string {
	identifier, input := readNextBits(input, 1)
	if identifier == "0" {
		return readBitLengthOperator(packet, input)
	}
	if identifier == "1" {
		input = readPacketLengthOperator(packet, input)
	}
	return input
}

func readPacketLengthOperator(packet *BitsPacket, input string) string {
	lengthStr, input := readNextBits(input, 11)
	packetsCount := bitsToInt(lengthStr)
	for i := 0; i < packetsCount; i++ {
		subPacket := BitsPacket{}
		input = readPacket(&subPacket, input)
		packet.SubPackets = append(packet.SubPackets, subPacket)
	}
	return input
}

func readBitLengthOperator(packet *BitsPacket, input string) string {
	lengthStr, input := readNextBits(input, 15)
	subPackets, input := readNextBits(input, bitsToInt(lengthStr))
	for len(subPackets) > 0 {
		subPacket := BitsPacket{}
		subPackets = readPacket(&subPacket, subPackets)
		packet.SubPackets = append(packet.SubPackets, subPacket)
	}
	return input
}

func readLiteralContents(packet *BitsPacket, input string) string {
	var nextGroupIdentifier, bits string
	for len(input) > 0 {
		nextGroupIdentifier, input = readNextBits(input, 1)
		bits, input = readNextBits(input, 4)
		packet.Contents += bits
		if nextGroupIdentifier == "0" {
			return input
		}
	}
	return input
}

func convertToBitsString(input string) string {
	hexList := strings.Split(input, "")
	bitsList := make([]string, len(hexList))
	for i, hex := range hexList {
		value, _ := strconv.ParseInt(hex, 16, 0)
		bitsList[i] = fmt.Sprintf("%04b", value)
	}
	return strings.Join(bitsList, "")
}

func readNextBits(input string, length int) (string, string) {
	bits := input[0:length]
	remaining := input[length:]
	return bits, remaining
}

func bitsToInt(input string) int {
	value, _ := strconv.ParseInt(input, 2, 0)
	return int(value)
}
