package hull

import (
	"testing"
)

//func TestInsertPolymerPairs(t *testing.T) {
//	type args struct {
//		sequence string
//		factory  PolymerFactory
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{"Step 1", args{sequence: "NNCB", factory: exampleFactory()}, "NCNBCHB"},
//		{"Step 2", args{sequence: "NCNBCHB", factory: exampleFactory()}, "NBCCNBBBCBHCB"},
//		{"Step 3", args{sequence: "NBCCNBBBCBHCB", factory: exampleFactory()}, "NBBBCNCCNBBNBNBBCHBHHBCHB"},
//		{"Step 4", args{sequence: "NBBBCNCCNBBNBNBBCHBHHBCHB", factory: exampleFactory()}, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := InsertPolymerPairs(tt.args.sequence, tt.args.factory); got != tt.want {
//				t.Errorf("InsertPolymerPairs() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestPolymerSequenceInfo(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		sequence := SequencePolymers("NNCB", exampleFactory(), 10)
		got := sequence.SequenceInfo()
		want := 1588
		if got != want {
			t.Errorf("TestPolymerSequenceInfo() = %v, want %v", got, want)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		sequence := SequencePolymers("NNCB", exampleFactory(), 40)
		got := sequence.SequenceInfo()
		want := 2188189693529
		if got != want {
			t.Errorf("TestPolymerSequenceInfo() = %v, want %v", got, want)
		}
	})
}

//NN
//NCN
//NBCCN
//NBBBCNCCN
//CBHCBHHNH
//
func exampleFactory() PolymerFactory {
	return BuildPolymerFactory([]string{
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	})
}
