package main

import (
	"fmt"
)

func main() {
	c := 0
	fmt.Scan(&c)

	tests := make([]string, c)
	for i := 0; i < c; i++ {
		fmt.Scan(&tests[i])
	}

	for _, test := range tests {
		lastchar := rune(0)
		count := 0

		for _, char := range test {
			if lastchar == char {
				count++
			} else if lastchar != 0 {
				if count < 2 {
					if count == 0 {
						fmt.Print(string(lastchar))
					} else {
						fmt.Printf("%v%v", string(lastchar), string(lastchar))
					}
				} else {
					fmt.Printf("%v%v", string(lastchar), count+1)
				}
				count = 0
			}
			lastchar = char
		}

		if count < 2 {
			if count == 0 {
				fmt.Print(string(lastchar))
			} else {
				fmt.Printf("%v%v", string(lastchar), string(lastchar))
			}
		} else {
			fmt.Printf("%v%v", string(lastchar), count+1)
		}

		fmt.Println()
	}
}
