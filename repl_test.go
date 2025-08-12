package main

import (
	"fmt"
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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " hElLo   JoE  aNd JaCk      ",
			expected: []string{"hello", "joe", "and", "jack"},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		failed := false
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("length of result does not match what is expected")
			failed = true
			failCount++
		} else {
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				// Check each word in the slice
				// if they don't match, use t.Errorf to print an error message
				// and fail the test
				if word != expectedWord {
					t.Errorf("result does not match expected")
					failed = true
					failCount++
					break
				}
			}
		}
		if !failed {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, c.input, c.expected, actual)
		} else {
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, c.input, c.expected, actual)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
