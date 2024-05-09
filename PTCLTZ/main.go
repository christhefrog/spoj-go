package main

import (
	"fmt"
)

/*
x0=s,
xn+1=3*xn+1, jeśli xn jest nieparzyste i
xn+1=xn/2, jeśli xn jest parzyste
*/

func collatz(s int) int {
	r, n := s, 0

	for r != 1 {
		if r%2 == 0 {
			r = r / 2
		} else {
			r = 3*r + 1
		}
		n++
	}

	return n
}

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		s := 0
		fmt.Scan(&s)

		fmt.Println(collatz(s))
	}

	return
}
