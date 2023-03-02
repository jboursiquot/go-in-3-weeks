package stringutils_test

import (
	"os"
	"testing"

	"github.com/jboursiquot/go-in-3-weeks/04-testing/stringutils"
)

// Exercise 1 solution
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

// Exercise 2 solution
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

// Exercise 3 solution
type testCase struct {
	input string
	want  string
}

var tests map[string]testCase

func TestMain(m *testing.M) {
	// setup
	tests = map[string]testCase{
		"basic":    {input: "HELLO", want: "hello"},
		"accented": {input: "CAFÉ", want: "café"},
	}
	code := m.Run()
	// teardown
	os.Exit(code)
}

func TestLowerWithTestCases(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := stringutils.Lower(tc.input)
			if tc.want != got {
				t.Fatalf("wanted %s but got %s", tc.want, got)
			}
		})
	}
}

// Exercise 4 solution

func benchUpper(str string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringutils.Upper(str)
	}
}

func BenchmarkUpper(b *testing.B) {
	benchUpper("hello", b)
}

func benchLower(str string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		stringutils.Lower(str)
	}
}

func BenchmarkLower(b *testing.B) {
	benchLower("HELLO", b)
}
