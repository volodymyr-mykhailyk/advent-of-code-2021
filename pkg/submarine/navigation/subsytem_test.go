package navigation

import "testing"

func TestSyntaxErrorScore(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"[]", 0},
		{"([])", 0},
		{"(((", 0},
		{"(]", 57},
		{"{([(<{}[<>[]}>{[]{[(<()>", 1197},
		{"[[<[([]))<([[{}[[()]]]", 3},
		{"[{[{({}]{}}([{[{{{}}([]", 57},
		{"[<(<(<(<{}))><([]([]()", 3},
		{"<{([([[(<>()){}]>(<<{{", 25137},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := syntaxErrorScore(tt.input); got != tt.want {
				t.Errorf("syntaxErrorScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoCompleteScore(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"[]", 0},
		{"([])", 0},
		{"(((", 31},
		{"[({(<(())[]>[[{[]{<()<>>", 288957},
		{"[(()[<>])]({[<{<<[]>>(", 5566},
		{"(((({<>}<{<{<>}{[]{[]{}", 1480781},
		{"{<[[]]>}<{[{[{[]{()[[[]", 995444},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 294},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := autoCompleteScore(tt.input); got != tt.want {
				t.Errorf("autoCompleteScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemErrorScore(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := exampleInput()
		got := SystemErrorScore(input)
		want := 26397
		if got != want {
			t.Errorf("TestSystemErrorScore() = %v, want %v", got, want)
		}
	})
}

func TestSystemAutocompleteScore(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		input := exampleInput()
		got := SystemAutocompleteScore(input)
		want := 288957
		if got != want {
			t.Errorf("TestSystemAutocompleteScore() = %v, want %v", got, want)
		}
	})
}

func exampleInput() []string {
	return []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
}
