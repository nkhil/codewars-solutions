/*
From: https://www.codewars.com/kata/56453a12fcee9a6c4700009c/go
Create a function close_compare that accepts 3 parameters: a, b, and an optional margin. The function should return whether a is lower than, close to, or higher than b.

a is considered "close to" b if margin is greater than or equal to the distance between a and b.

Please note the following:

When a is close to b, return 0.
Otherwise...

When a is less than b, return -1.

When a is greater than b, return 1.

If margin is not given, treat it as zero.

Assume: margin >= 0

Tip: Some languages have a way to make parameters optional.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(
		compareWithinMargin(8.1, 5.0, 3.0),
	)
}

func compareWithinMargin(a, b, margin float64) int {
	difference := math.Abs(a - b)
	isCloseTo := margin >= difference

	if isCloseTo {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}

/*
	Example solution
	func CloseCompare(a, b, margin float64) int {
		if math.Abs(a-b) <= margin { return 0 }
		if a > b { return 1 }
		return -1
	}
*/
