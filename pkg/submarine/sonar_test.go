package submarine

import "testing"

func TestExampleFloorIncreaseSpeed(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		readings := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := FloorIncreaseSpeed(readings)
		want := 7
		if got != want {
			t.Errorf("FloorIncreaseSpeed() = %v, want %v", got, want)
		}
	})
	t.Run("Euqal measurements", func(t *testing.T) {
		readings := []int{199, 199}
		got := FloorIncreaseSpeed(readings)
		want := 0
		if got != want {
			t.Errorf("FloorIncreaseSpeed() = %v, want %v", got, want)
		}
	})
}
