package navigation

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/structures"
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"testing"
)

func TestFindLowestRiskPath(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		plot := examplePathFinderInput()
		riskPlot := BuildRiskLevelPlot(plot, plot.Start(), plot.End())
		got := LowestRiskPathLevel(riskPlot, plot.End())
		want := 40
		if got != want {
			t.Errorf("TestFindLowestRiskPath() = %v, want %v", got, want)
		}
	})
	t.Run("Example Small", func(t *testing.T) {
		plot := examplePathFinderInputSmall()
		riskPlot := BuildRiskLevelPlot(plot, plot.Start(), structures.Point{X: 2, Y: 0})
		got := LowestRiskPathLevel(riskPlot, structures.Point{X: 2, Y: 0})
		want := 7
		if got != want {
			t.Errorf("TestFindLowestRiskPath() = %v, want %v", got, want)
		}
	})
	t.Run("Example Big", func(t *testing.T) {
		plot := examplePathFinderInput()
		plot = ScaleRiskPlot(plot, 5)
		//fmt.Printf("Map \n%v\n", plot.Presentation(""))
		riskPlot := BuildRiskLevelPlot(plot, plot.Start(), plot.End())
		//fmt.Printf("Risk \n%v\n", riskPlot.Presentation("\t\t"))
		got := LowestRiskPathLevel(riskPlot, plot.End())
		want := 315
		if got != want {
			t.Errorf("TestFindLowestRiskPath() = %v, want %v", got, want)
		}
	})
}

func examplePathFinderInput() structures.FlatValuePlot {
	input := []string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	}
	risks := transformations.SplitLines(input, "")
	return structures.FlatValuesFromString(risks)
}

func examplePathFinderInputSmall() structures.FlatValuePlot {
	input := []string{
		"181",
		"132",
		"111",
	}
	risks := transformations.SplitLines(input, "")
	return structures.FlatValuesFromString(risks)
}
