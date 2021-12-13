package vision

import (
	"testing"
)

func TestVisibleActivationDots(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		sheet := ActivationSheet(exampleInput())
		got := VisibleActivationDots(sheet)
		want := 18
		if got != want {
			t.Errorf("TestVisibleActivationDots() = %v, want %v", got, want)
		}
	})
	t.Run("After Y fold", func(t *testing.T) {
		sheet := ActivationSheet(exampleInput())
		sheet = FoldActivationSheet(sheet, "fold along y=7")
		got := VisibleActivationDots(sheet)
		want := 17
		if got != want {
			t.Errorf("TestVisibleActivationDots() = %v, want %v", got, want)
		}
	})
	t.Run("After both folds", func(t *testing.T) {
		sheet := ActivationSheet(exampleInput())
		sheet = FoldActivationSheet(sheet, "fold along y=7")
		sheet = FoldActivationSheet(sheet, "fold along x=5")
		got := VisibleActivationDots(sheet)
		want := 16
		if got != want {
			t.Errorf("TestVisibleActivationDots() = %v, want %v", got, want)
		}
	})
}
func exampleInput() []string {
	return []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
	}
}
