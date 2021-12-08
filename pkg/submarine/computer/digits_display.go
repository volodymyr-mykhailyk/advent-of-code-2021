package computer

import "strings"

var SignalsMap = map[string]int {
	"abcefg": 0,
	"cf": 1,
	"acdeg": 2,
	"acdfg": 3,
	"bcdf": 4,
	"abdfg": 5,
	"abdefg": 6,
	"acf": 7,
	"abcdefg": 8,
	"abcdfg": 9,
}

var LengthMap = map[int]int {
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

type Entry struct {
	SignalPattern []string
	DigitOutput []string
}

func DecodeDigits(entry Entry) []int {
	digits := make([]int, len(entry.DigitOutput))
	digits = decodeByLength(entry, digits)
	return digits
}

func DisplayOutput(entries []Entry) [][]int {
	digitsList := make([][]int, len(entries))
	for i, entry := range entries {
		digitsList[i] = DecodeDigits(entry)
	}
	return digitsList
}

func decodeByLength(entry Entry, digits []int) []int {
	for i, digitString := range entry.DigitOutput {
		if digit, ok := LengthMap[len(digitString)]; ok {
			digits[i] = digit
		} else {
			digits[i] = -1
		}
	}
	return digits
}

func ParseEntry(input string) Entry {
	entries := strings.Split(input, " | ")
	signals := strings.Split(entries[0], " ")
	digits := strings.Split(entries[1], " ")
	return Entry{SignalPattern: signals, DigitOutput: digits}
}