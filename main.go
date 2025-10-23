package main

import "fmt"

// FibSum returns the sum of all even Fibonacci numbers up to n.
func FibSum(n int) int {
	it1, it2 := 0, 1
	total := 0
	for it1 <= n {
		if it1&1 == 0 {
			total += it1
		}
		it1, it2 = it2, it1+it2
	}
	return total
}

func main() {
	cases := []int{-2, -1, 0, 1, 2, 8, 10, 34, 100}
	for _, n := range cases {
		fmt.Printf("FibSum(%d) = %d\n", n, FibSum(n))
	}
}
