package navigation

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"[]", true},
		{"([])", true},
		{"{()()()}", true},
		{"<([{}])>", true},
		{"[<>({}){}[([])<>]]", true},
		{"(((((((((())))))))))", true},
		{"(]", false},
		{"{()()()>", false},
		{"(((()))}", false},
		{"<([]){()}[{}])", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsValid(tt.input); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := SyntaxErrorScore(tt.input); got != tt.want {
				t.Errorf("SyntaxErrorScore() = %v, want %v", got, tt.want)
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
