package science

import "testing"

func TestSumElements(t *testing.T) {
	t.Run("Example ", func(t *testing.T) {
		input := []int{199, 200, 208}
		got := SumElements(input)
		want := 607
		if got != want {
			t.Errorf("TestSumElements() = %v, want %v", got, want)
		}
	})
}
