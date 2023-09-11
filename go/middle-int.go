// https://www.codewars.com/kata/545a4c5a61aa4c6916000755/train/go
// you need to create a function that when provided with a triplet, returns the index of the numerical element that lies between the other two elements.

// The input to the function will be an array of three distinct numbers (Haskell: a tuple).

// For example:

// gimme([2, 3, 1]) => 0
package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) int {
	sortedCopy := make([]int, 3)
	copy(sortedCopy, array[:])

	sort.Ints(sortedCopy)
	middleInt := sortedCopy[1]
	var result int

	for idx, el := range array {
		if el == middleInt {
			result = idx
			break
		}
	}

	return result
}

func main() {
	triplet := [3]int{-15, -10, 14}

	index := Gimme(triplet)

	fmt.Println("The index of the middle element is", index)
}

// From codewars
//
// func Gimme(array [3]int) int {
// 	arrayCopy := array

// 	arrSlice := arrayCopy[:]
// 	sort.Ints(arrSlice)
// 	val := arrSlice[1]

// 	for i, v := range array {
// 		if v == val {
// 			return i
// 		}
// 	}

// 	return 0
// }
