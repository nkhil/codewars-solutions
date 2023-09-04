// For https://www.codewars.com/kata/62a611067274990047f431a8/train/go

package main

import (
	"fmt"
)

func main() {
	result := Alternate(0, "blue", "red")

	fmt.Println(result)
}

func Alternate(n int, firstValue string, secondValue string) []string {
	mySlice := []string{}
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			mySlice = append(mySlice, firstValue)
		} else {
			mySlice = append(mySlice, secondValue)
		}
	}

	return mySlice
}
