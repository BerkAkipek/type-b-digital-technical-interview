package main

import "fmt"

// FibSum returns the sum of all even integers in given interval
func FibSum(n int) int {
	if n <= 1 {
		return 1
	}

	it1, it2 := 0, 1
	total := 0
	for it1 <= n {
		if it1&1 == 0 {
			total = total + it1
		}
		it1, it2 = it2, it1+it2
	}
	return total
}

func main() {
	fmt.Println(FibSum(2))
	fmt.Println(FibSum(8))
}
