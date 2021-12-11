package ocean

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"reflect"
	"testing"
)

func TestSimulateStep(t *testing.T) {
	t.Run("Example Simulation", func(t *testing.T) {
		plot := smallExample()
		SimulateStep(plot)
		got := plot.Printable("")
		want := []string{
			"34543",
			"40004",
			"50005",
			"40004",
			"34543",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSimulateStep() = %v, want %v", got, want)
		}
	})
	t.Run("Example Octopus Counts", func(t *testing.T) {
		plot := smallExample()
		got := SimulateStep(plot)
		want := 9
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSimulateStep() = %v, want %v", got, want)
		}
	})
}

func TestSimulateColony(t *testing.T) {
	t.Run("Example Simulation", func(t *testing.T) {
		plot := bigExample()
		got := SimulateColony(plot, 100)
		want := 1656
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSimulateColony() = %v, want %v", got, want)
		}
	})
}

func smallExample() structures.FlatValuePlot {
	data := []string{
		"11111",
		"19991",
		"19191",
		"19991",
		"11111",
	}
	return structures.FlatValuesFromString(transformations.SplitLines(data, ""))
}

func bigExample() structures.FlatValuePlot {
	data := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
	return structures.FlatValuesFromString(transformations.SplitLines(data, ""))
}
