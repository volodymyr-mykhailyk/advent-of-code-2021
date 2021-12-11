package transformations

import (
	"reflect"
	"testing"
)

func TestParseIntegers(t *testing.T) {
	content := ParseIntegers([]string{"1", "0", "-1000"})
	t.Run("Parsing", func(t *testing.T) {
		got := content
		want := []int{1, 0, -1000}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseIntegers()) = %v, want %v", got, want)
		}
	})
}

func TestParseBits(t *testing.T) {
	content := ParseBits([]string{"1", "10", "0101"})
	t.Run("Parsing", func(t *testing.T) {
		got := content
		want := []int{1, 2, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ParseBits()) = %v, want %v", got, want)
		}
	})
}

func TestSplitLines(t *testing.T) {
	content := SplitLines([]string{"11211", "12324", "11"}, "2")
	t.Run("Parsing", func(t *testing.T) {
		got := content
		want := [][]string{{"11", "11"}, {"1", "3", "4"}, {"11"}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSplitLines()) = %v, want %v", got, want)
		}
	})
}

func TestGroupOver(t *testing.T) {
	content := GroupOver([]string{"11", "22", "33", "22", "33", "11"}, "33")
	t.Run("Parsing", func(t *testing.T) {
		got := content
		want := [][]string{{"11", "22"}, {"22"}, {"11"}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestSplitLines()) = %v, want %v", got, want)
		}
	})
}
