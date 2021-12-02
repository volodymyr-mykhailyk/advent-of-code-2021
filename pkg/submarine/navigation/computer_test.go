package navigation

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"testing"
)

func TestPredictLinearPosition(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		commands := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
		got := PredictLinearPosition(commands, submarine.ZeroCoordinates())
		want := submarine.BuildCoordinates(15, 10, 0)
		if got != want {
			t.Errorf("TestPredictLinearPosition() = %v, want %v", got, want)
		}
	})
}

func TestPredictAimedPosition(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		commands := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
		got := PredictAimedPosition(commands, submarine.ZeroCoordinates())
		want := submarine.BuildCoordinates(15, 60, 10)
		if got != want {
			t.Errorf("TestPredictAimedPosition() = %v, want %v", got, want)
		}
	})
}
