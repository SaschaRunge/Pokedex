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
			input:    "  ",
			expected: []string{""},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HEllO  woRLd  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "Charmander,Bulbasaur /PIKACHU",
			expected: []string{"charmander,bulbasaur", "/pikachu"},
		},
		{
			input:    "72502    2552,1",
			expected: []string{"72502", "2552,1"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Differing slice lengths in cleanInput: actual %v, expected %v", len(actual), len(c.expected))
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Output mismatch: actual %s, expected %s", word, expectedWord)
			}
		}
	}
}
