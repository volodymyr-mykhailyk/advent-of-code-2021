package computer

import (
	"sort"
	"strconv"
	"strings"
)

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

func DecodeDigits(entry Entry) string {
	charsMap := buildDigitsMap(entry.SignalPattern)
	return decodeDigits(entry.DigitOutput, charsMap)
}

func buildDigitsMap(patterns []string) map[string]int {
	baseDigits := getBaseDigits(patterns)
	sixLength := filterByLength(patterns, 6)
	baseDigits[0] = getMatchingDigit(sixLength, []string{baseDigits[7]}, []string{baseDigits[4]})
	baseDigits[9] = getMatchingDigit(sixLength, []string{baseDigits[4]}, []string{baseDigits[0]})
	baseDigits[6] = getMatchingDigit(sixLength, []string{}, []string{baseDigits[1]})
	lowerSegment := findLowerSegment(baseDigits[6], baseDigits[1])
	fiveLength := filterByLength(patterns, 5)
	baseDigits[3] = getMatchingDigit(fiveLength, []string{baseDigits[1]}, []string{})
	baseDigits[5] = getMatchingDigit(fiveLength, []string{lowerSegment}, []string{baseDigits[3]})
	baseDigits[2] = getMatchingDigit(fiveLength, []string{}, []string{lowerSegment})
	return reverseMap(baseDigits)
}

func findLowerSegment(six string, one string) string {
	for _, char := range one {
		if matchingDigit(six, []string{string(char)}) {
			return string(char)
		}
	}
	return ""
}

func filterByLength(patterns []string, length int) []string {
	var result []string
	for _, pattern := range patterns {
		if len(pattern) == length {
			result = append(result, pattern)
		}
	}
	return result
}

func getMatchingDigit(patterns []string, matchers []string, nonMatchers []string) string {
	for _, pattern := range patterns {
		if matchingDigit(pattern, matchers) && (len(nonMatchers) == 0 || !matchingDigit(pattern, nonMatchers)) {
			return pattern
		}
	}
	return ""
}

func matchingDigit(pattern string, matchers []string) bool {
	for _, matcher := range matchers {
		for _, char := range matcher {
			if !strings.ContainsAny(pattern, string(char)) {
				return false
			}
		}
	}
	return true
}

func getBaseDigits(patterns []string) map[int]string {
	digits := make(map[int]string)
	for _, digitString := range patterns {
		if digit, ok := LengthMap[len(digitString)]; ok {
			digits[digit] = digitString
		}
	}
	return digits
}

func DisplayOutput(entries []Entry) []string {
	digitsList := make([]string, len(entries))
	for i, entry := range entries {
		digitsList[i] = DecodeDigits(entry)
	}
	return digitsList
}

func decodeDigits(digits []string, charMap map[string]int) string {
	result := ""
	for _, digitString := range digits {
		if digit, ok := charMap[digitString]; ok {
			result += strconv.FormatInt(int64(digit), 10)
		}
	}
	return result
}

func reverseMap(m map[int]string) map[string]int {
	n := make(map[string]int, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func ParseEntry(input string) Entry {
	entries := strings.Split(input, " | ")
	signals := sortSignals(strings.Split(entries[0], " "))
	digits := sortSignals(strings.Split(entries[1], " "))
	return Entry{SignalPattern: signals, DigitOutput: digits}
}

func sortSignals(signals []string) []string {
	for i, signal := range signals {
		s := strings.Split(signal, "")
		sort.Strings(s)
		signals[i] = strings.Join(s, "")
	}
	return signals
}