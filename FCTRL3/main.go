package main

import (
	"fmt"
)

func main() {
	d := 0
	fmt.Scan(&d)

	tests := make([]int, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&tests[i])
	}

	var factorials [10]string
	factorials[0] = "0 1"
	factorials[1] = "0 1"
	factorials[2] = "0 2"
	factorials[3] = "0 6"
	factorials[4] = "2 4"
	factorials[5] = "2 0"
	factorials[6] = "2 0"
	factorials[7] = "4 0"
	factorials[8] = "2 0"
	factorials[9] = "8 0"

	for _, test := range tests {
		if test > 9 {
			fmt.Println("0 0")
		} else {
			fmt.Println(factorials[test])
		}
	}
}
