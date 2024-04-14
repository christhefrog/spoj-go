package main

import (
	"fmt"
	"strconv"
)

func main() {
	d := 0
	fmt.Scan(&d)

	powers_periodicity := [10][]int{
		{1},
		{1},
		{6, 2, 4, 8},
		{1, 3, 9, 7},
		{6, 4},
		{5},
		{6},
		{1, 7, 9, 3},
		{6, 8, 4, 2},
		{1, 9},
	}

	tests := make([][2]int, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&tests[i][0])
		fmt.Scan(&tests[i][1])
	}

	for _, test := range tests {
		if test[0] == 1 {
			fmt.Println("1")
			continue
		} else if test[0]%10 == 0 {
			fmt.Println("0")
			continue
		}

		str := fmt.Sprint(test[0])
		last, _ := strconv.Atoi(str[len(str)-1:])

		fmt.Println(powers_periodicity[last][test[1]%len(powers_periodicity[last])])
	}

}
