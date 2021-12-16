package computer

import (
	"fmt"
	"math"
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
	value := 0
	switch packet.Type {
	case "000":
		for _, subPacket := range packet.SubPackets {
			value += subPacket.Value()
		}
	case "001":
		value = 1
		for _, subPacket := range packet.SubPackets {
			value *= subPacket.Value()
		}
	case "010":
		value = math.MaxInt
		for _, subPacket := range packet.SubPackets {
			if value > subPacket.Value() {
				value = subPacket.Value()
			}
		}
	case "011":
		for _, subPacket := range packet.SubPackets {
			if value < subPacket.Value() {
				value = subPacket.Value()
			}
		}
	case "100":
		value += bitsToInt(packet.Contents)
	case "101":
		if packet.SubPackets[0].Value() > packet.SubPackets[1].Value() {
			value = 1
		} else {
			value = 0
		}
	case "110":
		if packet.SubPackets[0].Value() < packet.SubPackets[1].Value() {
			value = 1
		} else {
			value = 0
		}
	case "111":
		if packet.SubPackets[0].Value() == packet.SubPackets[1].Value() {
			value = 1
		} else {
			value = 0
		}
	}
	return value
}

func bitsToInt(input string) int {
	value, _ := strconv.ParseInt(input, 2, 0)
	return int(value)
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
	readSubPackets(packet)
	input = removeTrailingZeros(input)
	return input
}

func removeTrailingZeros(input string) string {
	return input
}

func readSubPackets(packet *BitsPacket) {

}

func readPacketContents(packet *BitsPacket, input string) string {
	if packet.Type == "100" {
		return readLiteralContents(packet, input)
	} else {
		return readOperatorContents(packet, input)
	}
}

func readOperatorContents(packet *BitsPacket, input string) string {
	var identifier, lengthStr, subPackets string
	identifier, input = readNextBits(input, 1)
	if identifier == "0" {
		lengthStr, input = readNextBits(input, 15)
		length, _ := strconv.ParseInt(lengthStr, 2, 0)
		subPackets, input = readNextBits(input, int(length))
		for len(subPackets) > 0 {
			subPacket := BitsPacket{}
			subPackets = readPacket(&subPacket, subPackets)
			packet.SubPackets = append(packet.SubPackets, subPacket)
		}
	}
	if identifier == "1" {
		lengthStr, input = readNextBits(input, 11)
		length, _ := strconv.ParseInt(lengthStr, 2, 0)
		for i := 0; i < int(length); i++ {
			subPacket := BitsPacket{}
			input = readPacket(&subPacket, input)
			packet.SubPackets = append(packet.SubPackets, subPacket)
		}
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

func readPacketInfo(packet *BitsPacket, input string) string {
	packet.Version = input[0:3]
	packet.Type = input[3:6]
	return input[6:]
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
	remaining := "" + input[length:]
	return bits, remaining
}
