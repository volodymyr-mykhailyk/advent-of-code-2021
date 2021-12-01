package tasks

import (
	"reflect"
	"testing"
)

func TestReadLines(t *testing.T) {
	content := ReadLines("LICENSE")
	t.Run("Length", func(t *testing.T) {
		got := len(content)
		want := 25
		if got != want {
			t.Errorf("len(ReadLines()) = %v, want %v", got, want)
		}
	})
	t.Run("First Line", func(t *testing.T) {
		got := content[0]
		want := "This is free and unencumbered software released into the public domain."
		if got != want {
			t.Errorf("ReadLines()[0] = %v, want %v", got, want)
		}
	})
	t.Run("Empty Line", func(t *testing.T) {
		got := content[1]
		want := ""
		if got != want {
			t.Errorf("ReadLines()[1] = %v, want %v", got, want)
		}
	})
}

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
