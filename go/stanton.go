/*
https://www.codewars.com/kata/59a1cdde9f922b83ee00003b/train/go
The Stanton measure of an array is computed as follows: count the number of occurences for value 1 in the array. Let this count be n. The Stanton measure is the number of times that n appears in the array.

Write a function which takes an integer array and returns its Stanton measure.

Examples
The Stanton measure of [1, 4, 3, 2, 1, 2, 3, 2] is 3, because 1 occurs 2 times in the array and 2 occurs 3 times.

The Stanton measure of [1, 4, 1, 2, 11, 2, 3, 1] is 1, because 1 occurs 3 times in the array and 3 occurs 1 time.
*/

package main

import "fmt"

func main() {
	var arr = []int{1, 4, 1, 2, 11, 2, 3, 1}
	fmt.Println(
		StantonMeasure(arr),
	)
}

func StantonMeasure(arr []int) int {
	var numOfOnes int
  for _, el := range arr { 
		if el == 1 { numOfOnes++ }
	}

	var count int

	for _, elem := range arr {
		if elem == numOfOnes { count++ }
	}

	return count
}

/*
Solution from Codewars:

func StantonMeasure(arr []int) int {
  seen := make(map[int]int, len(arr))
  for _, v := range arr {
    seen[v]++
  }
  return seen[seen[1]]
}
*/
