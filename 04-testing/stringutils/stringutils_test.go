package stringutils_test

import (
	"testing"

	"github.com/jboursiquot/go-next-steps/04-testing/stringutils"
)

func TestUpper(t *testing.T) {
	input := "hello"
	want := "HELLO"
	got := stringutils.Upper(input)
	if want != got {
		t.Fatalf("wanted %s but got %s", want, got)
	}

	input = "café"
	want = "CAFÉ"
	got = stringutils.Upper(input)
	if want != got {
		t.Fatalf("wanted %s but got %s", want, got)
	}
}

// TestLower uses table-driven style of testing.
func TestLower(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"basic":    {input: "HELLO", want: "hello"},
		"accented": {input: "CAFÉ", want: "café"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := stringutils.Lower(tc.input)
			if tc.want != got {
				t.Fatalf("wanted %s but got %s", tc.want, got)
			}
		})
	}
}
