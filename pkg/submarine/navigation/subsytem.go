package navigation

import (
	"math"
	"sort"
	"strings"
)

var subsystemGroups = []string{
	"()",
	"[]",
	"{}",
	"<>",
}

var autoCompleteScores = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

var invalidPatterns = map[string]int{
	"(]": 57,
	"(}": 1197,
	"(>": 25137,
	"[)": 3,
	"[}": 1197,
	"[>": 25137,
	"{)": 3,
	"{]": 57,
	"{>": 25137,
	"<)": 3,
	"<]": 57,
	"<}": 1197,
}

func SystemErrorScore(input []string) int {
	score := 0
	for _, line := range input {
		score += syntaxErrorScore(line)
	}
	return score
}

func SystemAutocompleteScore(input []string) int {
	var scores []int
	for _, line := range input {
		if !isError(line) {
			scores = append(scores, autoCompleteScore(line))
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func autoCompleteScore(line string) int {
	normalized := reduceValidGroups(line)
	missingCharacters := autoCompleteString(normalized)
	score := 0
	for _, char := range missingCharacters {
		score *= 5
		score += autoCompleteScores[char]
	}
	return score
}

func autoCompleteString(normalized string) []string {
	chars := strings.Split(normalized, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}

func syntaxErrorScore(input string) int {
	corruptedGroup := corruptedGroup(reduceValidGroups(input))
	if corruptedGroup == "" {
		return 0
	} else {
		return invalidPatterns[corruptedGroup]
	}
}

func isError(input string) bool {
	normalized := reduceValidGroups(input)
	for pattern := range invalidPatterns {
		if strings.Contains(normalized, pattern) {
			return true
		}
	}
	return false
}

func corruptedGroup(input string) string {
	indexMap := make(map[string]int)
	for group := range invalidPatterns {
		index := strings.Index(input, group)
		if index != -1 {
			indexMap[group] = index
		}
	}
	if len(indexMap) > 0 {
		firstGroupIndex := math.MaxInt
		firstGroup := ""
		for group, index := range indexMap {
			if index < firstGroupIndex {
				firstGroupIndex = index
				firstGroup = group
			}
		}
		return firstGroup
	} else {
		return ""
	}
}

func reduceValidGroups(input string) string {
	for hasGroup(input) {
		for _, group := range subsystemGroups {
			input = strings.ReplaceAll(input, group, "")
		}
	}
	return input
}

func hasGroup(input string) bool {
	for _, group := range subsystemGroups {
		if strings.Contains(input, group) {
			return true
		}
	}
	return false
}
