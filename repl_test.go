package main

import (
	"testing"
)

// TABLE DRIVEN TESTING - The Gold Standard
func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "How many   words can you split",
			expected: []string{"how", "many", "words", "can", "you", "split"},
		},
		{
			input:    " th1s  wun  IZ different",
			expected: []string{"th1s", "wun", "iz", "different"},
		},
		{
			input:    "  two  Words   ",
			expected: []string{"two", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check length of slice vs. expected
		if len(actual) != len(c.expected) {
			// error and continue
			t.Errorf("For Input %q: expected length %d, got %d (actual slice: %v)", c.input, len(c.expected), len(actual), actual)
			continue // skip checking if lengths dont match

		} // if lengths match, check each word
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				//use t.Errorf to print an error message & FAIL the test
				t.Errorf("For input %q: at index %d, expected %q, got %q", c.input, i, expectedWord, word)
			}
		}
	}
}

/*
	type testCase struct {

		// input parameters
		value float64
		base string
		conv string
		date string

		// expected values
		expected float
	}
*/
