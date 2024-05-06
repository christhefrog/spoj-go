package main

import (
	"fmt"
	"math"
)

func main() {
	t := 0
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		n, cookies, sum := 0, 0, 0
		fmt.Scan(&n, &cookies)

		for j := 0; j < n; j++ {
			time := 0.0
			fmt.Scan(&time)
			sum += int(math.Floor(86400.0 / time))
		}

		fmt.Println(math.Ceil(float64(sum) / float64(cookies)))
	}
}
