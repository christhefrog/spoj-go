package main

import (
	"fmt"
)

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		n, first := 0, 0
		fmt.Scan(&n, &first)
		for j := 1; j < n; j++ {
			el := 0
			fmt.Scan(&el)
			fmt.Print(el, " ")
		}
		fmt.Print(first, " ")
	}
}
