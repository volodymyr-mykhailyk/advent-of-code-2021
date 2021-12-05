package sonar

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/ocean"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/tasks"
	"reflect"
	"testing"
)

func TestBuildVentMap(t *testing.T) {
	t.Run("Horizontal", func(t *testing.T) {
		readings := buildVents([]string{"1,1 -> 3,1"})
		got := BuildVentMap(readings)
		want := ocean.VentMap{ocean.Point{X: 1, Y: 1}: 1, ocean.Point{X: 2, Y: 1}: 1, ocean.Point{X: 3, Y: 1}: 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestBuildVentMap() = %v, want %v", got, want)
		}
	})
	t.Run("Overlapping", func(t *testing.T) {
		readings := buildVents([]string{"1,1 -> 2,1", "2,0 -> 2,1"})
		got := BuildVentMap(readings)
		want := ocean.VentMap{ocean.Point{X: 1, Y: 1}: 1, ocean.Point{X: 2, Y: 1}: 2, ocean.Point{X: 2, Y: 0}: 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestBuildVentMap() = %v, want %v", got, want)
		}
	})
	t.Run("Vertical", func(t *testing.T) {
		readings := buildVents([]string{"1,1 -> 1,3"})
		got := BuildVentMap(readings)
		want := ocean.VentMap{ocean.Point{X: 1, Y: 1}: 1, ocean.Point{X: 1, Y: 2}: 1, ocean.Point{X: 1, Y: 3}: 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestBuildVentMap() = %v, want %v", got, want)
		}
	})
}

func buildVents(inputs []string) []ocean.Vent {
	vents := make([]ocean.Vent, len(inputs))
	tasks.Iterate(inputs, func(input string, i int) {
		vents[i] = ocean.VentFromString(input)
	})
	return vents
}

