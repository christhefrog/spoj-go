package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func main() {
	t := 0
	fmt.Scan(&t)

	tests := make([][2]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&tests[i][0])
		fmt.Scan(&tests[i][1])
	}

	for _, test := range tests {
		fmt.Println(gcd(test[0], test[1]))
	}
}
