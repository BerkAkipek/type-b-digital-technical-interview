package main

import "testing"

func TestFibSum(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"input is -2", -2, 0},
		{"input is 0", 0, 0},
		{"input is 1", 1, 0},
		{"input is 2", 2, 2},
		{"input is 3", 3, 2},
		{"input is 8", 8, 10},
		{"input is 10", 10, 10},
		{"input is 34", 34, 44},
		{"input is 100", 100, 44},
		{"input is 1000", 1000, 798},
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
