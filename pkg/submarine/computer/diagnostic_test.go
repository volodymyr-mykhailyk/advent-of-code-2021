package computer

import (
	"testing"
)

func TestPowerConsumption(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		reports := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
		got := PowerConsumption(reports)
		want := 198
		if got != want {
			t.Errorf("TestPowerConsumption() = %v, want %v", got, want)
		}
	})
}
