package hull

import (
	"math"
	"strings"
)

type PolymerFactory map[string]string

func SequencePolymers(sequence string, factory PolymerFactory, steps int) string {
	for i := 0; i < steps; i++ {
		sequence = InsertPolymerPairs(sequence, factory)
	}
	return sequence
}

func InsertPolymerPairs(sequence string, factory PolymerFactory) string {
	if len(sequence) > 1 {
		return insertPair(sequence[:2], factory) + InsertPolymerPairs(sequence[1:], factory)
	} else {
		return sequence
	}
}

func PolymerSequenceInfo(sequence string) int {
	elements := elementsCount(sequence)
	uniq, common := elementPairs(elements)
	return common - uniq
}

func elementPairs(elements map[string]int) (int, int) {
	min := math.MaxInt
	max := 0
	for _, count := range elements {
		if min > count {
			min = count
		}
		if max < count {
			max = count
		}
	}
	return min, max
}

func elementsCount(sequence string) map[string]int {
	elements := make(map[string]int)
	for _, element := range strings.Split(sequence, "") {
		elements[element]++
	}
	return elements
}

func insertPair(pair string, factory PolymerFactory) string {
	if element, ok := factory[pair]; ok {
		return pair[:1] + element
	} else {
		return pair
	}
}

func BuildPolymerFactory(input []string) PolymerFactory {
	factory := make(PolymerFactory, len(input))
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		factory[parts[0]] = parts[1]
	}
	return factory
}
