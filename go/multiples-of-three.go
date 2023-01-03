/*
https://www.codewars.com/kata/61123a6f2446320021db987d/train/go
Given a positive integer n: 0 < n < 1e6,
remove the last digit until you're left with a number that is a multiple of three.

Return n if the input is already a multiple of three,
and if no such number exists, return null, a similar empty value, or -1.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		PrevMultOfThree(952406),
	)
}

func PrevMultOfThree(n int) interface{} {
  // write your code here
  // your function should return an int or a nil
	if IsDivisibleByThree(n) {
		return n
	}

	for {
		if (n < 10) && !IsDivisibleByThree(n) {
			return nil
		}

		n = n / 10

		if IsDivisibleByThree(n) {
			return n
		}
	}
}

func IsDivisibleByThree(n int) bool {
	return (n % 3) == 0
}

/*
Better solution from CodeWars

func PrevMultOfThree(n int) interface{} {
  
	for i := n; i > 0; i /= 10 {
		if i % 3 == 0 {
				return i
		}
	}

    return nil
}

*/