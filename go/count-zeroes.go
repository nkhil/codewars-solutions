// https://www.codewars.com/kata/52f787eb172a8b4ae1000a34/train/go
// This is by Bard
// This algorithm works by repeatedly shifting the number n
// to the right by 1 bit and checking if the least significant
// bit is not set. If the least significant bit is not set,
// the algorithm increments the number of zeroes.

// The function CountZeroes() takes an integer as input
// and returns an integer. The integer returned by the
// function is the number of bits that are equal to zero
// in thebinary representation of the input number.

// The example you gave has the binary representation 10011010010.
// The number of zeroes in this binary representation is 2.
// So, the function should return 2 in this case.

package main

import (
	"fmt"
)

func CountZeroes(n int) int {
	// Initialize the number of zeroes to 0.
	count := 0

	// Loop while n is not zero.
	for n != 0 {
		// Check if the least significant bit is not set.
		if n&1 == 0 {
			// Increment the number of zeroes.
			count++
		}

		// Shift n to the right by 1 bit.
		n >>= 1
	}

	// Return the number of zeroes.
	return count
}

func main() {
	n := 1234

	count := CountZeroes(n)

	fmt.Println("The number of zeroes in", n, "is", count)
}
