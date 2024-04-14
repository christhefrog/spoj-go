package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	max := int(math.Sqrt(float64(n))) + 1

	if n == 1 {
		return false
	}

	for i := 2; i < max; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 0
	fmt.Scan(&n)

	tests := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tests[i])
	}

	for _, test := range tests {
		if isPrime(test) {
			fmt.Println("TAK")
		} else {
			fmt.Println("NIE")
		}
	}
}
