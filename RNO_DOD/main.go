package main

import (
	"fmt"
)

func sum(arr []int) int {
	ans := 0
	for _, number := range arr {
		ans += number
	}

	return ans
}

func main() {
	t, n := 0, 0
	fmt.Scan(&t)

	tests := make([][]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		tests[i] = make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&tests[i][j])
		}
	}

	for _, test := range tests {
		fmt.Println(sum(test))
	}
}
