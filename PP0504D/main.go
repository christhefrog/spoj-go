package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func reverse(array []byte) []byte {
	a, b, c := 0, len(array)-1, byte(0)

	for b > a {
		c = array[a]
		array[a] = array[b]
		array[b] = c
		a++
		b--
	}

	return array
}

func main() {
	d := 0
	fmt.Scan(&d)

	tests := make([]float32, d)
	for i := 0; i < d; i++ {
		fmt.Scan(&tests[i])
	}

	for _, test := range tests {
		var buf bytes.Buffer
		err := binary.Write(&buf, binary.LittleEndian, test)
		if err != nil {
		}

		arr := reverse(buf.Bytes())

		for _, hex := range arr {
			fmt.Printf("%x ", hex)
		}

		fmt.Println()
	}

}
