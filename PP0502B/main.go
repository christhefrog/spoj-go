package main

import (
	"fmt"
)

func reverse(array []int) []int {
	a, b, c := 0, len(array)-1, 0

	for b > a {
		c = array[a]
		array[a] = array[b]
		array[b] = c
		a++
		b--
	}

	return array
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
		reversed := reverse(test)

		for _, v := range reversed {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

}
