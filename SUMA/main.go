package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum += num
		fmt.Println(sum)
	}
}
