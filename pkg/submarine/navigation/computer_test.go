package navigation

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"testing"
)

func TestPredictPosition(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		commands := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
		got := PredictPosition(commands, submarine.Coordinates{})
		want := submarine.BuildCoordinates(15, 10)
		if got != want {
			t.Errorf("TestPredictPosition() = %v, want %v", got, want)
		}
	})
}
