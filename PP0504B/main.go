package main

import (
	"fmt"
)

func main() {
	d := 0
	fmt.Scan(&d)

	tests := make([][2]string, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&tests[i][0])
		fmt.Scan(&tests[i][1])
	}

	for _, test := range tests {
		ile := 0
		if len(test[0]) < len(test[1]) {
			ile = len(test[0])
		} else {
			ile = len(test[1])
		}

		for i := 0; i < ile; i++ {
			fmt.Print(string(test[0][i]))
			fmt.Print(string(test[1][i]))
		}

		fmt.Println()
	}

}
