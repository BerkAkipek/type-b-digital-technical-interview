package main

import "testing"

func TestFibSum(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"n is -2", -2, 0},
		{"n is 0", 0, 0},
		{"n is 1", 1, 0},
		{"n is 2", 2, 2},
		{"n is 3", 3, 2},
		{"n is 8", 8, 10},
		{"n is 10", 10, 10},
		{"n is 34", 34, 44},
		{"n is 100", 100, 44},
		{"n is 1000", 1000, 798},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := FibSum(testCase.input)
			if result != testCase.expected {
				t.Errorf("FibSum(%d) = %d; want %d", testCase.input, result, testCase.expected)
			}
		})
	}
}
