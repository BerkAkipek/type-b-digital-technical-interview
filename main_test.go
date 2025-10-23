package main

import "testing"

func TestFibSum(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n is 0", 0, 1},
		{"n is 1", 1, 1},
		{"n is 2", 2, 2},
		{"n is 3", 3, 2},
		{"n is 8", 8, 10},
		{"n is 10", 10, 10},
		{"n is 34", 34, 44},
		{"n is 100", 100, 44},
		{"n is 1000", 1000, 798},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FibSum(tt.n)
			if got != tt.want {
				t.Errorf("FibSum(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}
