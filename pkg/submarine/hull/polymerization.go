package hull

import (
	"math"
	"strconv"
	"strings"
)

type Element struct {
	symbol string
	count  int
}
type PolymerFactory map[string]string
type PolymerElements map[string]int

func (elements PolymerElements) Merge(other PolymerElements) {
	for k, v := range other {
		elements[k] += v
	}
}

func (elements PolymerElements) SequenceInfo() int {
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
	return max - min
}

type SequenceCache map[string]PolymerElements

func SequencePolymers(seed string, factory PolymerFactory, steps int) PolymerElements {
	elements := make(PolymerElements)
	cache := make(SequenceCache)
	for i, element := range strings.Split(seed, "") {
		elements[element]++
		if i < len(seed)-1 {
			element1 := seed[i : i+1]
			element2 := seed[i+1 : i+2]
			elements.Merge(sequencePair(factory, cache, element1, element2, steps))
		}
	}
	return elements
}

func sequencePair(factory PolymerFactory, cache SequenceCache, element1 string, element2 string, steps int) PolymerElements {
	if steps > 0 {
		cacheKey := element1 + element2 + strconv.Itoa(steps)
		if elements, ok := cache[cacheKey]; ok {
			return elements
		} else {
			newElement := factory[element1+element2]
			elements := make(PolymerElements)
			elements[newElement]++
			elements.Merge(sequencePair(factory, cache, element1, newElement, steps-1))
			elements.Merge(sequencePair(factory, cache, newElement, element2, steps-1))
			cache[cacheKey] = elements
			return elements
		}
	} else {
		return PolymerElements{}
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
