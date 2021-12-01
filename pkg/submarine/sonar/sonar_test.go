package sonar

import "testing"

func TestLinearIncreaseSpeed(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		readings := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := LinearIncreaseSpeed(readings)
		want := 7
		if got != want {
			t.Errorf("LinearIncreaseSpeed() = %v, want %v", got, want)
		}
	})
	t.Run("Euqal measurements", func(t *testing.T) {
		readings := []int{199, 199}
		got := LinearIncreaseSpeed(readings)
		want := 0
		if got != want {
			t.Errorf("LinearIncreaseSpeed() = %v, want %v", got, want)
		}
	})
}
func TestProgressiveIncreaseSpeed(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		readings := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := AveragedIncreaseSpeed(readings)
		want := 5
		if got != want {
			t.Errorf("AveragedIncreaseSpeed() = %v, want %v", got, want)
		}
	})
	t.Run("Euqal measurements", func(t *testing.T) {
		readings := []int{199, 199}
		got := LinearIncreaseSpeed(readings)
		want := 0
		if got != want {
			t.Errorf("AveragedIncreaseSpeed() = %v, want %v", got, want)
		}
	})
}
