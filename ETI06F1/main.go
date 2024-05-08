package main

import (
	"fmt"
	"math"
)

const PI = 3.141592654

func main() {
	r, d := 0.0, 0.0
	fmt.Scan(&r, &d)

	a := d / 2.0

	//a^2 + x^2 = r^2
	//x^2 = r^2 - d^2
	radius := math.Sqrt(r*r - a*a)

	// PI * r^2
	area := PI * radius * radius

	fmt.Printf("%.2f", area)
}
