package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"reflect"
	"testing"
)

func TestLowPoints(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings(exampleInput())
		got := LowPoints(heightMap)
		want := []ocean.Point{{1, 0}, {9, 0}, {2, 2}, {6, 4}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
	t.Run("Example equal points", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings([]string{"12", "12"})
		got := LowPoints(heightMap)
		var want []ocean.Point
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
}
func TestLargestBasins(t *testing.T) {
	//t.Run("Example 1", func(t *testing.T) {
	//	heightMap := ocean.HeightMapFromStrings(exampleInput())
	//	got := LowPoints(heightMap)
	//	want := []ocean.Point{{1, 0}, {9, 0}, {2, 2}, {6, 4}}
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("TestLowPoints() = %v, want %v", got, want)
	//	}
	//})
	t.Run("Example small points", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings([]string{"191", "229"})
		got := LargestBasins(heightMap, LowPoints(heightMap))
		want := [][]ocean.Point{{{0, 0}, {0, 1}, {1, 1}}, {{2, 0}}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLargestBasins() = %v, want %v", got, want)
		}
	})
}

func TestLowPointsRiskLevel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings(exampleInput())
		got := LowPointsRiskLevel(heightMap, LowPoints(heightMap))
		want := 15
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestRiskLevel() = %v, want %v", got, want)
		}
	})
}

func TestBasinsRiskLevel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings(exampleInput())
		lowPoints := LowPoints(heightMap)
		got := BasinsRiskLevel(heightMap, LargestBasins(heightMap, lowPoints))
		want := 1134
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestBasinsRiskLevel() = %v, want %v", got, want)
		}
	})
}

func exampleInput() []string {
	return []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
}
