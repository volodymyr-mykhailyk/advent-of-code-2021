package computer

import (
	"reflect"
	"testing"
)

func TestDecodeDigits(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		got := DecodeDigits(ParseEntry("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"))
		want := []int{8, -1, -1, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestDecodeDigits() = %v, want %v", got, want)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		got := DecodeDigits(ParseEntry("edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"))
		want := []int{-1, 7, 8, 1}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("TestDecodeDigits() = %v, want %v", got, want)
		}
	})
}

func exampleEntries() []Entry {
	inputs := exampleInput()
	entries := make([]Entry, len(inputs))
	for i, input := range inputs {
		entries[i] = ParseEntry(input)
	}
	return entries
}
func exampleInput() []string {
	return []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
}

