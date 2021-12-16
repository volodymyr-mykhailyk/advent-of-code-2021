package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"reflect"
	"testing"
)

func TestLowPoints(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := exampleSmokeBasin()
		got := LowPoints(heightMap)
		want := []structures.Point{{1, 0}, {9, 0}, {2, 2}, {6, 4}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
	t.Run("Example equal points", func(t *testing.T) {
		heightMap := structures.FlatValuesFromString(transformations.SplitLines([]string{"12", "12"}, ""))
		got := LowPoints(heightMap)
		var want []structures.Point
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
}
func TestLargestBasins(t *testing.T) {
	t.Run("Example small points", func(t *testing.T) {
		heightMap := structures.FlatValuesFromString(transformations.SplitLines([]string{"191", "229"}, ""))
		got := LargestBasins(heightMap, LowPoints(heightMap))
		want := [][]structures.Point{{{0, 0}, {0, 1}, {1, 1}}, {{2, 0}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLargestBasins() = %v, want %v", got, want)
		}
	})
}

func TestLowPointsRiskLevel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := exampleSmokeBasin()
		got := LowPointsRiskLevel(heightMap, LowPoints(heightMap))
		want := 15
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestRiskLevel() = %v, want %v", got, want)
		}
	})
}

func TestBasinsRiskLevel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := exampleSmokeBasin()
		lowPoints := LowPoints(heightMap)
		got := BasinsRiskLevel(heightMap, LargestBasins(heightMap, lowPoints))
		want := 1134
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestBasinsRiskLevel() = %v, want %v", got, want)
		}
	})
}

func exampleSmokeBasin() structures.FlatValuePlot {
	lines := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	return structures.FlatValuesFromString(transformations.SplitLines(lines, ""))
}
