package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " THIS is A FrEaKiNg Nightare  !  ",
			expected: []string{"this", "is", "a", "freaking", "nightare", "!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// 1. Check length first
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input %q: got %v, want %v", c.input, len(actual), len(c.expected))
			continue // Skip element check if lengths don't match
		}

		// 2. Check each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word mismatch at index %d: got %q, want %q", i, actual[i], c.expected[i])
			}
		}
	}
}
