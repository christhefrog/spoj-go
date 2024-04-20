package main

import (
	"fmt"
	"math"
	"sort"
)

func distance(a int, b int) float64 {
	return math.Sqrt(math.Pow(-float64(a), 2) + math.Pow(-float64(b), 2))
}

func main() {
	t := 0
	fmt.Scan(&t)

	tests := make([]map[string][]int, t)
	for i := 0; i < t; i++ {
		n := 0
		fmt.Scan(&n)
		tests[i] = make(map[string][]int)

		for j := 0; j < n; j++ {
			name := ""
			fmt.Scan(&name)

			tests[i][name] = make([]int, 2)
			fmt.Scan(&tests[i][name][0])
			fmt.Scan(&tests[i][name][1])
		}
	}

	for _, test := range tests {
		distances := make(map[string]float64)

		for name, point := range test {
			distances[name] = distance(point[0], point[1])
		}

		keys := make([]string, 0, len(distances))
		for key := range distances {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return distances[keys[i]] < distances[keys[j]]
		})

		for _, ans := range keys {
			fmt.Println(ans, test[ans][0], test[ans][1])
		}
	}
}
