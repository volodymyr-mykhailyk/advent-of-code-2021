package navigation

import (
	"math"
	"strings"
)

var subsystemGroups = []string{
	"()",
	"[]",
	"{}",
	"<>",
}

var syntaxScores = map[string]int{
	")": 3,
	"]": 57,
	"{": 1197,
	">": 25137,
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

func IsValid(input string) bool {
	if reduceValidGroups(input) != "" {
		return false
	} else {
		return true
	}
}

func SystemErrorScore(input []string) int {
	score := 0
	for _, line := range input {
		score += SyntaxErrorScore(line)
	}
	return score
}

func SyntaxErrorScore(input string) int {
	corruptedGroup := corruptedGroup(reduceValidGroups(input))
	if corruptedGroup == "" {
		return 0
	} else {
		return invalidPatterns[corruptedGroup]
	}
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
