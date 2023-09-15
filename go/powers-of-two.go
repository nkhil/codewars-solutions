package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) []uint64 {
	result := []uint64{}
	for i := 0; i <= n; i++ {
		result = append(result, uint64(math.Pow(2, float64(i))))
	}
	return result
}

func main() {
	fmt.Println(
		PowersOfTwo(0),
		PowersOfTwo(4),
	)
}
