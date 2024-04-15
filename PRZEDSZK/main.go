package main

import (
	"fmt"
)

func lcm(a int, b int) int {
	x, y := a, b

	for x != y {
		if x > y {
			y += b
		} else {
			x += a
		}
	}

	return x
}

func main() {
	n := 0
	fmt.Scan(&n)

	tests := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tests[i][0])
		fmt.Scan(&tests[i][1])
	}

	for _, test := range tests {
		fmt.Println(lcm(test[0], test[1]))
	}
}
