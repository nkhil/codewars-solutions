/*
https://www.codewars.com/kata/598d91785d4ce3ec4f000018/train/go

Given a string "abc" and assuming that each letter in the string
has a value equal to its position in the alphabet, our string will
have a value of 1 + 2 + 3 = 6.
This means that: a = 1, b = 2, c = 3 ....z = 26.

You will be given a list of strings and your task will be to
return the values of the strings as explained above multiplied
by the position of that string in the list. For our purpose,
position begins with 1.

nameValue ["abc","abc abc"] should return [6,24] because of [ 6 * 1, 12 * 2 ]. Note how spaces are ignored.

"abc" has a value of 6, while "abc abc" has a value of 12. Now, the value at position 1 is multiplied by 1 while the value at position 2 is multiplied by 2.

Input will only contain lowercase characters and spaces.

Good luck!
*/

package main

import (
	"fmt"
)
func main() {
	fmt.Println(
		CalculateArrayScore([]string{"codewars","abc","xyz"}),
	)
}

var scores = [256]byte{
	' ': 0, // ignore spaces
	'A': 1, 'a': 1,
	'B': 2, 'b': 2,
	'C': 3, 'c': 3,
	'D': 4, 'd': 4,
	'E': 5, 'e': 5,
	'F': 6, 'f': 6,
	'G': 7, 'g': 7,
	'H': 8, 'h': 8,
	'I': 9, 'i': 9,
	'J': 10, 'j': 10,
	'K': 11, 'k': 11,
	'L': 12, 'l': 12,
	'M': 13, 'm': 13,
	'N': 14, 'n': 14,
	'O': 15, 'o': 15,
	'P': 16, 'p': 16,
	'Q': 17, 'q': 17,
	'R': 18, 'r': 18,
	'S': 19, 's': 19,
	'T': 20, 't': 20,
	'U': 21, 'u': 21,
	'V': 22, 'v': 22,
	'W': 23, 'w': 23,
	'X': 24, 'x': 24,
	'Y': 25, 'y': 25,
	'Z': 25, 'z': 26,
}

func CalculateWordScore(str string) int {
	var score int
	for _, c := range str {
		score += int(scores[c])
	}
	return score
}

func CalculateArrayScore(my_list []string) []int {
	result := []int{}
	for i, s := range my_list {
		result = append(result, (CalculateWordScore(s) * (i + 1)))
	}
	return result
}

/* Better solution from CodeWars

func NameValue(my_list []string) []int {
	var result = make([]int, len(my_list))

	for idx, str := range my_list {
		for _, chr := range str {
			if chr >= 'a' && chr <= 'z' {
				result[idx] += int(chr-'a') + 1
			}
		}

		result[idx] = result[idx] * (idx + 1)
	}

	return result
}

*/