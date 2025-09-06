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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "                          CHARMANDER			",
			expected: []string{"charmander"},
		},
		{
			input:    "PoKeMoN      ",
			expected: []string{"pokemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length actual: %v		Length expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word: %s	Expected word: %s", word, expectedWord)
			}
		}
	}
}
