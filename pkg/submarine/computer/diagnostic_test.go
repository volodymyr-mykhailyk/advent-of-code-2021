package computer

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/transformations"
	"testing"
)

func TestPowerConsumption(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		reports := transformations.ParseBits([]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"})
		got := PowerConsumption(reports, 5)
		want := 198
		if got != want {
			t.Errorf("TestPowerConsumption() = %v, want %v", got, want)
		}
	})
}

func TestLifeSupportRating(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		reports := transformations.ParseBits([]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"})
		got := LifeSupportRating(reports, 5)
		want := 230
		if got != want {
			t.Errorf("TestLifeSupportRating() = %v, want %v", got, want)
		}
	})
}
