package transformations

import (
	"log"
	"strconv"
	"strings"
)

func SplitLines(input []string, separator string) [][]string {
	result := make([][]string, len(input))
	for i, s := range input {
		result[i] = strings.Split(s, separator)
	}
	return result
}

func ParseIntegers(input []string) []int {
	result := make([]int, len(input))
	for i, v := range input {
		parsed, err := strconv.Atoi(v)
		result[i] = parsed
		if err != nil {
			log.Fatalf("unable to parse int: %v", v)
		}
	}
	return result
}

func ParseBits(input []string) []int {
	var result []int
	for _, v := range input {
		parsed, err := strconv.ParseInt(v, 2, 0)
		result = append(result, int(parsed))
		if err != nil {
			log.Fatalf("unable to parse int: %v", v)
		}
	}
	return result
}


