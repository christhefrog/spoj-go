package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	num := fmt.Sprint(n)
	a, b := 0, len(num)-1

	for b > a {
		if num[a] != num[b] {
			return false
		}
		a++
		b--
	}

	return true
}

func reverse(n int) int {
	num := fmt.Sprint(n)
	reversed := ""

	for i := len(num) - 1; i >= 0; i-- {
		reversed += string(num[i])
	}

	res, _ := strconv.Atoi(string(reversed))
	return res
}

func main() {
	t := 0
	fmt.Scan(&t)

	tests := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&tests[i])
	}

	for _, test := range tests {
		i := 0
		for !isPalindrome(test) {
			test = test + reverse(test)
			i++
		}

		fmt.Println(test, i)
	}
}
