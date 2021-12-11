package computer

import (
	"reflect"
	"testing"
)

func TestDecodeDigits(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := DecodeDigits(ParseEntry("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"))
		want := "8394"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestDecodeDigits() = %v, want %v", got, want)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		got := DecodeDigits(ParseEntry("edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"))
		want := "9781"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestDecodeDigits() = %v, want %v", got, want)
		}
	})
	t.Run("Example 3", func(t *testing.T) {
		got := DecodeDigits(ParseEntry("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"))
		want := "5353"
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestDecodeDigits() = %v, want %v", got, want)
		}
	})
}
