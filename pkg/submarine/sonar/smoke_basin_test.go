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
		want := []int{1, 0, 5, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
	t.Run("Example equal points", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings([]string{"12", "12"})
		got := LowPoints(heightMap)
		var want []int
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestLowPoints() = %v, want %v", got, want)
		}
	})
}

func TestRiskLevel(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		heightMap := ocean.HeightMapFromStrings(exampleInput())
		got := RiskLevel(LowPoints(heightMap))
		want := 15
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestRiskLevel() = %v, want %v", got, want)
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
